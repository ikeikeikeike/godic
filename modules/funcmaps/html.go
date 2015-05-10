package funcmaps

import (
	"fmt"
	"html"
	"html/template"
	"strings"
	"time"

	gq "github.com/PuerkitoBio/goquery"
	"github.com/ikeikeikeike/godic/modules/configs"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/ikeikeikeike/gopkg/str"
	"github.com/kennygrant/sanitize"
	"github.com/mattn/go-runewidth"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func ToAge(bd time.Time) int {
	if bd.UnixNano() < 0 {
		return 0
	}

	at := time.Now()
	age := at.Year() - bd.Year()
	if (at.Month() <= bd.Month()) && (at.Day() <= bd.Day()) {
		age -= 1
	}
	return age
}

func SafeHTML(text string) template.HTML {
	return template.HTML(text)
}

func EscapeHTML(text string) string {
	return html.EscapeString(text)
}

func SanitizeHTML(text string) string {
	content, _ := sanitize.HTMLAllowing(text)
	content = sanitize.HTML(text)
	return html.EscapeString(content)
}

func MarkdownHTML(markdown string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(markdown))
	contentHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(contentHtml)
}

func Nl2br(in string) string {
	return strings.Replace(in, "\n", "<br>", -1)
}

func DiffLineTypeToStr(diffType int) string {
	switch diffType {
	case 2:
		return "add"
	case 3:
		return "del"
	case 4:
		return "tag"
	}
	return "same"
}

func DiffTypeToStr(diffType int) string {
	diffTypes := map[int]string{
		1: "add", 2: "modify", 3: "del",
	}
	return diffTypes[diffType]
}

func Truncate(in string, length int) string {
	return runewidth.Truncate(in, length, "...")
}

// Expect sorted string-length.
func AutoLink(html string, names []string) string {
	doc, err := gq.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return html
	}

	for _, name := range names {
		if len([]rune(name)) <= 1 {
			continue
		}

		doc.Find("*").Each(func(i int, s *gq.Selection) {
			if strings.Contains(s.Text(), name) {
				t := s.Text()
				h, _ := s.Html()
				if t != "" && h == "" {
					html := strings.Replace(t, name, fmt.Sprintf(`<a href="/d/%[1]s">%[1]s</a>`, name), 1)
					s.ReplaceWithHtml(html)
				}
			}
		})
	}

	res, err := doc.Find("body").Html()
	if err != nil {
		return html
	}

	return res
}

type Img struct {
	*image.FileInfo

	Src string
	Alt string
}

func ExtractIMGs(html string) (imgs []*Img) {
	doc, err := gq.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}

	info := image.NewInfo()
	info.Header("User-Agent", configs.Settings.UserAgent)

	doc.Find("img").Each(func(i int, s *gq.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		alt, _ := s.Attr("alt")
		if alt == "" {
			alt, _ = s.Attr("title")
		}

		src = str.Clean(src)
		f, err := info.Fetch(src)
		if err == nil {
			img := &Img{FileInfo: f, Src: src, Alt: alt}
			imgs = append(imgs, img)
		}
	})

	return
}
