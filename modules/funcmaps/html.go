package funcmaps

import (
	"html"
	"html/template"
	"strings"
	"time"
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

func EscapeHTML(in string) string {
	return html.EscapeString(in)
}

func Nl2br(in string) string {
	return strings.Replace(in, "\n", "<br>", -1)
}
