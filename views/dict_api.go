package views

import (
	"fmt"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models/dict"
	"github.com/ikeikeikeike/godic/modules/git"
	"github.com/ikeikeikeike/godic/views/forms"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type APIResponse struct {
	ok  bool
	sha string
	msg string
}

func UpdateDict(params martini.Params, commit forms.Commit, errs binding.Errors, r render.Render) {
	log.Debugln("UpdateDict action !!!!!")

	if len(errs) > 0 {
		msg := fmt.Sprintf("valid error (%d):\n%+v", len(errs), errs)
		r.JSON(404, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	sha1, err := Repo.SaveFile(params["name"], commit.Content, commit.Message)
	if err != nil {
		msg := fmt.Sprintf("Update file error: %s", err)
		r.JSON(200, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	r.JSON(200, APIResponse{ok: true, sha: sha1.String()})
}

func CreateDict(params martini.Params, commit forms.Commit, errs binding.Errors, r render.Render) {
	log.Debugln("CreateDict action !!!!!")

	if len(errs) > 0 {
		msg := fmt.Sprintf("valid error (%d):\n%+v", len(errs), errs)
		r.JSON(404, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	m, _ := dict.FirstOrCreateByCommit(commit)

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	// Create
	sha1, err := Repo.SaveFile(m.Name, commit.Content, commit.Message)
	if err != nil {
		msg := fmt.Sprintf("Create file error: %s", err)
		r.JSON(200, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	r.JSON(200, APIResponse{ok: true, sha: sha1.String()})
}

func DeleteDict(params martini.Params, html html.HTMLContext, r render.Render) {
	log.Debugln("DeleteDict action !!!!!")

	if params["name"] == "" {
		r.JSON(404, APIResponse{ok: false, msg: "Invalid name"})
		return
	}

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}
