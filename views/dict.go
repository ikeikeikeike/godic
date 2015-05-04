package views

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/models/dict"
	"github.com/ikeikeikeike/godic/modules/git"
	git2go "github.com/libgit2/git2go"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var BasePath, RepoPath string

func init() {
	BasePath, _ = os.Getwd()
	RepoPath = path.Join(BasePath, "repo")
}

func DictIndex(r render.Render, html html.HTMLContext) {
	log.Debugln("IndexDict action !!!!!")

	var dicts []*models.Dict
	dict.Dicts().Find(&dicts)

	html["Dicts"] = dicts

	r.HTML(200, "dict/index", html)
}

func DictHistory(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("IndexDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	m := &models.Dict{}
	if err := dict.FirstByName(params["name"], m).Error; err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	history, err := repo.GetFileHistory(params["name"], 1)
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
	m := &models.Dict{}
	if err := dict.FirstByName(params["name"], m).Error; err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	diff, err := repo.GetDiffRange(params["fromsha1"], params["tosha1"], 0)
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	html["Name"] = params["name"]
	html["Diff"] = diff

	r.HTML(200, "dict/compare", html)
}

func ShowDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("ShowDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	m := &models.Dict{}
	if err := dict.FirstByName(params["name"], m).Error; err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	blob, err := repo.GetFileBlob(params["name"])
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	markdown := blob.Contents()
	unsafe := blackfriday.MarkdownCommon(markdown)
	contentHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	html["Name"] = params["name"]
	html["Yomi"] = m.Yomi
	html["Content"] = string(markdown)
	html["ContentHTML"] = string(contentHtml)

	html["Dict"] = m

	r.HTML(200, "dict/show", html)
}

func NewDict(r render.Render, params martini.Params, html html.HTMLContext, req *http.Request) {
	log.Debugln("NewDict action !!!!!")

	name := params["name"]
	if name == "" {
		name = "no_title"
	}

	html["Name"] = name
	html["Content"] = ""

	category := &models.Category{}
	models.DB.Where("prefix = ?", req.URL.Query().Get("category")).First(&category)
	if category.ID > 0 {
		html["Category"] = category
	}

	var categories []*models.Category
	models.DB.Find(&categories)
	html["Categories"] = categories

	bytes, err := ioutil.ReadFile(path.Join(BasePath, "template.txt"))
	if err == nil {
		html["Content"] = fmt.Sprintf(string(bytes), name)
	}

	r.HTML(200, "dict/edit", html)
}

func EditDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("EditDict action !!!!!")
	name := params["name"]

	if name == "" {
		r.HTML(404, "errors/404", html)
		return
	}
	m := &models.Dict{}
	if err := dict.FirstByName(name, m).Error; err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	blob, err := repo.GetFileBlob(name)
	if err != nil {
		r.HTML(404, "errors/404", html)
		return
	}

	l, err := repo.GetFileHistory(name, 1)
	if err != nil {
		log.Warn(err)
	}

	var c *git2go.Commit
	for e := l.Front(); e != nil; e = e.Next() {
		c = e.Value.(*git2go.Commit)
	}

	var categories []*models.Category
	models.DB.Find(&categories)

	html["Name"] = params["name"]
	html["Yomi"] = m.Yomi
	html["Content"] = string(blob.Contents())
	html["Categories"] = categories

	html["Dict"] = m
	html["Category"] = m.Category
	html["Commit"] = c

	r.HTML(200, "dict/edit", html)
}
