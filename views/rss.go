package views

import (
	"fmt"
	gohtml "html"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/models/dict"
	"github.com/ikeikeikeike/godic/modules/configs"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
	dichttp "github.com/ikeikeikeike/godic/modules/http"
	"github.com/martini-contrib/render"
)

func LatestRSS(r render.Render, html html.HTMLContext, routes martini.Routes, req *http.Request) {
	log.Debugln("LatestRss action !!!!!")

	var dicts []*models.Dict
	dict.RelationDB().Limit(10).Order("dicts.created_at DESC").
		Find(&dicts)

	rss := genFeeds(dicts, html, routes, req)

	r.Header().Set("Content-Type", "text/xml; charset=utf-8")
	r.Data(200, []byte(rss))
}

func ModifiedRSS(r render.Render, html html.HTMLContext, routes martini.Routes, req *http.Request) {
	log.Debugln("ModifiedRSS action !!!!!")

	var dicts []*models.Dict
	dict.RelationDB().Limit(10).Order("dicts.updated_at DESC").
		Find(&dicts)

	rss := genFeeds(dicts, html, routes, req)

	r.Header().Set("Content-Type", "text/xml; charset=utf-8")
	r.Data(200, []byte(rss))
}

func newFeed(h html.HTMLContext) *feeds.Feed {
	now := time.Now()
	settings := configs.Settings

	return &feeds.Feed{
		Title:       gohtml.EscapeString(settings.AppName),
		Link:        &feeds.Link{Href: h["Meta"].(*html.Meta).Host},
		Description: gohtml.EscapeString(settings.Description),
		Author:      &feeds.Author{settings.Author, settings.Email},
		Created:     now,
	}
}

func genFeeds(dicts []*models.Dict, html html.HTMLContext, routes martini.Routes, req *http.Request) string {
	feed := newFeed(html)

	for _, dict := range dicts {

		href := dichttp.BuildRequestUrl(req, routes.URLFor("show", dict.Name))
		body := funcmaps.Truncate(funcmaps.SanitizeHTML(funcmaps.MarkdownHTML(dict.Content)), 80)

		src := dichttp.BuildRequestUrl(req, "/static/img/siteicon/apple-touch-icon.png")
		if dict.Image != nil {
			src = dict.Image.Src
		}

		feed.Items = append(feed.Items, &feeds.Item{
			Title: dict.Name,
			Link:  &feeds.Link{Href: href},
			Description: fmt.Sprintf(`<![CDATA[
				<div>
					<a href="%s">
						<img src="%s" style="max-height:400px; max-width:300px">
						<div>
							<h6>%s</h6>
							<p>%s</p>
							<p>続きを読む</p>
						</div>
					</a>
				</div>
				]]>`,
				href, src, dict.Name, funcmaps.Nl2br(body),
			),
			Author:  &feeds.Author{"アノニマス", ""},
			Created: dict.CreatedAt,
		})
	}

	rss, _ := feed.ToRss()
	return gohtml.UnescapeString(rss)
}
