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

	App.Get("/", Home)

	App.Get("/index", Index)
	App.Get("/create", Create)
}
