package views

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
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
	App.Use(html.ProvideHTMLHeader)
	App.Use(html.ProvideHTMLMeta)
	// App.Use(html.ProvideMartiniParams)

	App.Get("/", Home).Name("root")

	App.Get("/d/index", html.SetParams, DictIndex).Name("index")
	App.Get("/d/history/:name", html.SetParams, DictHistory).Name("history")

	App.Get("/d/new/", NewDict).Name("new")
	App.Get("/d/new/:name", NewDict).Name("new")
	App.Get("/d/:name", html.SetParams, ShowDict).Name("show")
	App.Get("/d/edit/:name", html.SetParams, EditDict).Name("edit")

	App.Group("/_d", func(r martini.Router) {
		r.Get("/:name", func(r render.Render, p martini.Params) {
			r.Redirect("/d/" + p["name"])
		})
		r.Put("/:name", binding.Bind(Commit{}), UpdateDict).Name("api_put")
		r.Post("/:name", binding.Bind(Commit{}), CreateDict).Name("api_post")
		r.Delete("/:name", DeleteDict).Name("api_delete")
	})

	App.NotFound(func() (int, string) {
		return 404, "not found"
	})
}
