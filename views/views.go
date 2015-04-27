package views

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
)

var App *martini.ClassicMartini

func init() {
	App = martini.Classic()

	App.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
		Funcs:      funcmaps.HelperFuncs,
	}))

	App.Use(html.GenHTMLContext())
	App.Use(html.ProvideHTMLHeader)
	App.Use(html.ProvideHTMLMeta)

	App.Get("/", Home)

	App.Get("/d/index", DictIndex)
	App.Get("/d/history/:name", DictHistory)

	App.Get("/d/new/:name", NewDict)
	App.Get("/d/edit/:name", EditDict)

	App.Group("/_d", func(r martini.Router) {
		r.Get("/:name", NewDict)
		r.Put("/:name", binding.Bind(Commit{}), UpdateDict)
		r.Post("/:name", binding.Bind(Commit{}), CreateDict)
		r.Delete("/:name", DeleteDict)
	})
}
