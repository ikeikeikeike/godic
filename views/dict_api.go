package views

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/k0kubun/pp"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type APIResponse struct {
	ok  bool
	sha string
	msg string
}

type Commit struct {
	Name    string `form:"name" binding:"required"`
	Content string `form:"content" binding:"required"`
	Message string `form:"message"`
}

// func (p Post) Validate(errors binding.Errors, req *http.Request) binding.Errors { return errros }

func UpdateDict(params martini.Params, commit Commit, errs binding.Errors, r render.Render) {
	log.Println("UpdateDict action !!!!!")

	if params["name"] == "" {
		r.JSON(404, APIResponse{ok: false, msg: "Invalid name"})
		return
	}

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

func CreateDict(params martini.Params, commit Commit, errs binding.Errors, r render.Render) {
	log.Println("CreateDict action !!!!!")

	if params["name"] == "" {
		pp.Println("not found")
		r.JSON(404, APIResponse{ok: false, msg: "Not Found"})
		return
	}

	if len(errs) > 0 {
		pp.Println("errors")
		msg := fmt.Sprintf("valid error (%d):\n%+v", len(errs), errs)
		r.JSON(404, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	// if current_app.config.get('ALLOW_ANON') and current_user.is_anonymous():
	// return dict(error=True, message="Anonymous posting not allowed"), 403

	// if cname in current_app.config.get('WIKI_LOCKED_PAGES'):
	// return dict(error=True, message="Page is locked"), 403

	// Create
	sha1, err := Repo.SaveFile(params["name"], commit.Content, commit.Message)
	if err != nil {
		msg := fmt.Sprintf("Create file error: %s", err)
		r.JSON(200, APIResponse{ok: false, msg: msg})
		log.Fatal(msg)
		return
	}

	r.JSON(200, APIResponse{ok: true, sha: sha1.String()})
}

func DeleteDict(params martini.Params, html html.HTMLContext, r render.Render) {
	log.Println("DeleteDict action !!!!!")

	html["Name"] = params["name"]
	html["Content"] = ""

	r.HTML(200, "dict/edit", html)
}
