package views

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/models"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
)

func CreateComments(r render.Render, p martini.Params, s sessions.Session, u sessionauth.User, routes martini.Routes, comment models.Comment, errs binding.Errors) {
	log.Debugln("CreateComments action !!!!!")

	r.Redirect(routes.URLFor("show", p["name"]))

	if len(errs) > 0 {
		log.Warning(errs)
		s.AddFlash("必須項目を入力してください")
		return
	}

	if v, ok := u.(*models.User); ok && v.ID > 0 {
		comment.User = v
	}

	var d models.Dict
	models.DB.Where("dicts.name = ?", p["name"]).First(&d)

	c := &comment
	c.BeforeSave()  // Polymoficed save not working hook method.

	d.Comments = append(d.Comments, c)
	models.DB.Save(&d)
}
