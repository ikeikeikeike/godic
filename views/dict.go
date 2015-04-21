package views

import "github.com/martini-contrib/render"

func Index(r render.Render) {
	r.HTML(200, "dict/index", "")
}

func Create(r render.Render) {
	r.HTML(200, "dict/create", "")
}
