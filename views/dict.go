package views

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/modules/git"
	git2go "github.com/libgit2/git2go"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var Repo *git.Repo
var BasePath, RepoPath string

func init() {
	BasePath, _ = os.Getwd()
	RepoPath = path.Join(BasePath, "repo")

	Repo = git.NewRepo()
	Repo.Init(RepoPath)
}

func DictIndex(r render.Render, html html.HTMLContext) {
	log.Debugln("IndexDict action !!!!!")

	names, err := Repo.FolderFileNames()
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	html["Names"] = names

	r.HTML(200, "dict/index", html)
}

func DictHistory(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("IndexDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	history, err := Repo.GetFileHistory(params["name"], 1)
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	html["Name"] = params["name"]
	html["History"] = history

	r.HTML(200, "dict/history", html)
}

func CompareDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("CompareDict action !!!!!")

	if params["name"] == "" || params["fromsha1"] == "" || params["tosha1"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	diff, err := Repo.GetDiffRange(params["fromsha1"], params["tosha1"], 0)
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	html["Name"] = params["name"]
	html["Diff"] = diff

	r.HTML(200, "dict/compare", html)
}

func NewDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("NewDict action !!!!!")

	name := params["name"]
	if name == "" {
		name = "no_title"
	}

	html["Name"] = name
	html["Content"] = ""

	bytes, err := ioutil.ReadFile(path.Join(BasePath, "template.txt"))
	if err == nil {
		html["Content"] = fmt.Sprintf(string(bytes), name)
	}

	r.HTML(200, "dict/edit", html)
}

func ShowDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("ShowDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	blob, err := Repo.GetFileBlob(params["name"])
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	markdown := blob.Contents()
	unsafe := blackfriday.MarkdownCommon(markdown)
	contentHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	html["Name"] = params["name"]
	html["Content"] = string(markdown)
	html["ContentHTML"] = string(contentHtml)

	r.HTML(200, "dict/show", html)
}

func EditDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("EditDict action !!!!!")
	name := params["name"]

	if name == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	blob, err := Repo.GetFileBlob(name)
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}
	m := &models.Dict{}
	if err := models.DB.Where("name = ?", name).Preload("Image").Preload("Category").First(m).Error; err != nil {
		r.HTML(404, "errors/404", html)
		return
	}
	l, err := Repo.GetFileHistory(name, 1)
	if err != nil {
		log.Warn(err)
	}

	var c *git2go.Commit
	for e := l.Front(); e != nil; e = e.Next() {
		c = e.Value.(*git2go.Commit)
	}

	html["Name"] = params["name"]
	html["Yomi"] = m.Yomi
	html["Content"] = string(blob.Contents())

	html["Dict"] = m
	html["Commit"] = c

	r.HTML(200, "dict/edit", html)
}
