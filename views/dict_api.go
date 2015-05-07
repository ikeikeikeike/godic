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
	Ok  bool   `json:"ok"`
	Sha string `json:"sha"`
	Msg string `json:"msg"`
}

func UpdateDict(params martini.Params, commit forms.Commit, errs binding.Errors, r render.Render) {
	log.Debugln("UpdateDict action !!!!!")

	if len(errs) > 0 {
		msg := fmt.Sprintf("valid error (%d):\n%+v", len(errs), errs)
		r.JSON(400, APIResponse{Ok: false, Msg: msg})
		log.Warn(msg)
		return
	}

	m := dict.UpdateByCommit(commit)

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	sha1, err := repo.SaveFile(m.Name, commit.Content, commit.Message)
	if err != nil {
		msg := fmt.Sprintf("Update file error: %s", err)
		r.JSON(400, APIResponse{Ok: false, Msg: msg})
		log.Warn(msg)
		return
	}

	r.JSON(200, APIResponse{Ok: true, Sha: sha1.String()})
}

func CreateDict(params martini.Params, commit forms.Commit, errs binding.Errors, r render.Render) {
	log.Debugln("CreateDict action !!!!!")

	if len(errs) > 0 {
		msg := fmt.Sprintf("valid error (%d):\n%+v", len(errs), errs)
		r.JSON(400, APIResponse{Ok: false, Msg: msg}) // TODO: Now this error message is used in javascript alert, next time must be change to japanese.
		log.Warn(msg)
		return
	}

	m, created := dict.FirstOrCreateByCommit(commit)
	if !created {
		msg := fmt.Sprintf("%[1]sは既に存在します: > <a target='blank_' href='/d/%[1]s'>%[1]s</a>", m.Name)
		r.JSON(400, APIResponse{Ok: false, Msg: msg})
		log.Warn(msg)
		return
	}

	repo := git.NewRepo()
	repo.Init(path.Join(RepoPath, m.GetPrefix()))

	// Create
	sha1, err := repo.SaveFile(m.Name, commit.Content, commit.Message)
	if err != nil {
		msg := fmt.Sprintf("Create file error: %s", err)
		r.JSON(400, APIResponse{Ok: false, Msg: msg})
		log.Warn(msg)
		return
	}

	r.JSON(200, APIResponse{Ok: true, Sha: sha1.String()})
}

func DeleteDict(params martini.Params, html html.HTMLContext, r render.Render) {
	log.Debugln("DeleteDict action !!!!!")

	if params["name"] == "" {
		r.JSON(404, APIResponse{Ok: false, Msg: "Invalid name"})
		return
	}

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}
