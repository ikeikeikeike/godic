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
	"github.com/ikeikeikeike/godic/models/category"
	"github.com/ikeikeikeike/godic/models/dict"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
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

	all := category.CategoriesALL()
	for _, c := range all {
		c.DictLoader(25)
	}
	html["CategoriesALL"] = all

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

	html["Dict"] = m
	html["Name"] = params["name"]
	html["History"] = history

	r.HTML(200, "dict/history", html, render.HTMLOptions{"layout-editor"})
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

	html["Dict"] = m
	html["Name"] = params["name"]
	html["Diff"] = diff

	r.HTML(200, "dict/compare", html, render.HTMLOptions{"layout-editor"})
}

func ShowDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("ShowDict action !!!!!")

	if params["name"] == "" {
		r.HTML(404, "errors/404", html)
		return
	}

	html["Name"] = params["name"]

	var cdicts, udicts []*models.Dict
	dict.RelationDB().Order("dicts.created_at DESC").Limit(5).Find(&cdicts)
	dict.RelationDB().Order("dicts.created_at DESC").Limit(5).Find(&udicts)

	html["CreatedDicts"] = cdicts
	html["UpdatedDicts"] = udicts

	all := category.CategoriesALL()
	for _, c := range all {
		c.DictLoader(10)
	}
	html["CategoriesALL"] = all

	m := &models.Dict{}
	if err := dict.FirstByName(params["name"], m).Error; err != nil {
		r.HTML(200, "dict/notfound", html)
		return
	}

	html["Dict"] = m
	html["Yomi"] = m.Yomi
	html["Content"] = m.Content
	html["ContentHTML"] = m.ContentHTML

	if params["sha1"] != "" {
		repo := git.NewRepo()
		repo.Init(path.Join(RepoPath, m.GetPrefix()))

		blob, err := repo.GetFileBlobWithHash(params["name"], params["sha1"])
		if err != nil {
			r.HTML(200, "dict/notfound", html)
			return
		}

		markdown := blob.Contents()
		unsafe := blackfriday.MarkdownCommon(markdown)
		contentHtml := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		html["Content"] = string(markdown)
		html["ContentHTML"] = string(contentHtml)
	}

	r.HTML(200, "dict/show", html)
}

func NewDict(r render.Render, params martini.Params, html html.HTMLContext, req *http.Request) {
	log.Debugln("NewDict action !!!!!")

	name := params["name"]
	if name == "" {
		name = "no_title"
	}
	name = funcmaps.ToCanonical(name)

	html["Name"] = name
	html["Content"] = ""

	c := &models.Category{}
	models.DB.Where("prefix = ?", req.URL.Query().Get("category")).First(&c)
	if c.ID > 0 {
		html["Category"] = c
	}

	html["Categories"] = category.Categories()

	bytes, err := ioutil.ReadFile(path.Join(BasePath, "template.txt"))
	if err == nil {
		html["Content"] = fmt.Sprintf(string(bytes), name, req.URL.Query().Get("image"))
	}

	r.HTML(200, "dict/edit", html, render.HTMLOptions{"layout-editor"})
}

func EditDict(r render.Render, params martini.Params, html html.HTMLContext) {
	log.Debugln("EditDict action !!!!!")
	name := funcmaps.ToCanonical(params["name"])

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

	html["Name"] = params["name"]
	html["Yomi"] = m.Yomi
	html["Content"] = string(blob.Contents())
	html["Categories"] = category.Categories()

	html["Dict"] = m
	html["Category"] = m.Category
	html["Commit"] = c

	r.HTML(200, "dict/edit", html, render.HTMLOptions{"layout-editor"})
}
