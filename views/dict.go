package views

import (
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/modules/git"
	"github.com/martini-contrib/render"
)

var Repo *git.Repo

func init() {
	p, _ := os.Getwd()

	Repo = git.NewRepo()
	Repo.Init(path.Join(p, "repo"))
}

func DictIndex(r render.Render, html html.HTMLContext) {
	log.Println("IndexDict action !!!!!")

	names, err := Repo.FolderFileNames()
	if err != nil {
		r.HTML(404, "404 not found", html)
		return
	}

	html["Names"] = names

	r.HTML(200, "dict/index", html)
}

func DictHistory(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Println("IndexDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "404 not found", html)
		return
	}
	history, err := Repo.GetFileHistory(params["name"], 1)
	if err != nil {
		r.HTML(404, "404 not found", html)
		return
	}

	html["Name"] = params["name"]
	html["History"] = history

	r.HTML(200, "dict/history", html)
}

func NewDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Infoln("NewDict action !!!!!")

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}

func EditDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Println("EditDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "404 not found", html)
		return
	}
	blob, err := Repo.GetFileBlob(params["name"])
	if err != nil {
		r.HTML(404, "404 not found", html)
		return
	}

	html["Name"] = params["name"]
	html["Content"] = string(blob.Contents())

	r.HTML(200, "dict/edit", html)
}
