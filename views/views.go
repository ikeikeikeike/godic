package views

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	goauth2 "golang.org/x/oauth2"

	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
	"github.com/ikeikeikeike/godic/views/forms"
)

var App *martini.ClassicMartini

func init() {
	App = martini.Classic()

	funcs := append(funcmaps.HelperFuncs, template.FuncMap{"urlFor": App.URLFor})
	App.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
		Funcs:      funcs,
	}))

	App.Use(html.GenHTMLContext())
	App.Use(html.HTMLHeader)
	App.Use(html.HTMLMeta)
	App.Use(html.HTMLSettings)

	App.Use(sessions.Sessions(
		"godic_sesssion", sessions.NewCookieStore([]byte("secret09131ffl2"))),
	)

	App.Use(oauth2.Github(&goauth2.Config{
		ClientID:     "0.0",
		ClientSecret: "o.o",
		// RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes: []string{"user:email", "read:org"},
	}))
	App.Use(oauth2.Facebook(&goauth2.Config{
		ClientID:     "0.0",
		ClientSecret: "o.o",
		// RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes: []string{"user:email", "read:org"},
	}))
	App.Use(oauth2.Google(&goauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		// RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes: []string{"https://www.googleapis.com/auth/drive"},
	}))

	App.Get("/", func(r render.Render) { r.Redirect("/d/index") }).Name("root")
	App.Group("/abouts", func(r martini.Router) {
		r.Get("/sitemap", func(r render.Render) { r.Redirect("/") }).Name("abouts_sitemap")
	})

	App.Group("/d", func(r martini.Router) {
		r.Get("/index", DictIndex).Name("index")
		r.Get("/new/", NewDict).Name("new")
		r.Get("/new/:name", NewDict).Name("new")
		r.Get("/:name", ShowDict).Name("show")
		r.Get("/edit/:name", EditDict).Name("edit")
		r.Get("/history/:name", DictHistory).Name("history")
		r.Get(`/compare/:name/(?P<fromsha1>[^\.]+)\.{2,3}(?P<tosha1>.+)`, CompareDict).Name("compare")
	}, html.RequestParams)

	App.Group("/_d", func(r martini.Router) {
		r.Get("/:name", func(r render.Render, p martini.Params) { r.Redirect("/d/" + p["name"]) })
		r.Put("/:name", binding.Bind(forms.Commit{}), UpdateDict).Name("api_put")
		r.Post("/:name", binding.Bind(forms.Commit{}), CreateDict).Name("api_post")
		r.Delete("/:name", DeleteDict).Name("api_delete")
	})

	App.NotFound(func() (int, string) {
		return 404, "not found"
	})
}
