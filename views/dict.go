package views

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/martini-contrib/render"
)

func IndexDict(r render.Render, html html.HTMLContext) {
	log.Println("IndexDict action !!!!!")
	r.HTML(200, "dict/index", html)
}

func NewDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Infoln("NewDict action !!!!!")

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}

func EditDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Println("EditDict action !!!!!")

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}
