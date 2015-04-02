package views

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var App *martini.ClassicMartini

func init() {
	App = martini.Classic()

	App.Use(render.Renderer(render.Options{
		Extensions: []string{".html"},
	}))

	App.Get("/", Root)

	App.Get("/render", func(r render.Render) {
		r.HTML(200, "hello", "jeremy")
	})
}

func Root() string {
	return "Hello world!"
}
