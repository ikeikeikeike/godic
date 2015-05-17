package views

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
)

func LoginAccounts(r render.Render, html html.HTMLContext) {
	log.Debugln("Login action !!!!!")
	r.HTML(200, "accounts/login", html)
}

func SaveLoginAccounts(r render.Render, html html.HTMLContext, session sessions.Session, form models.User, errs binding.Errors, req *http.Request) {
	log.Debugln("SaveLogin action !!!!!")

	if len(errs) > 0 {
		log.Warning(errs)

		html["Errors"] = errs
		r.HTML(200, "accounts/login", html)
		return
	}

	var errors binding.Errors
	errors = append(errors, binding.Error{
		Message: "メールアドレスまたはパスワードが間違っています。",
	})

	user := models.User{}
	err := models.DB.Where("email = ? AND password = ?", form.Email, form.Password).First(&user).Error
	if err != nil {
		log.Warning(err)

		html["Errors"] = errors
		r.HTML(200, "accounts/login", html)
		return
	}

	err = sessionauth.AuthenticateSession(session, &user)
	if err != nil {
		log.Warning(err)

		html["Errors"] = errors
		r.HTML(200, "accounts/login", html)
		return
	}

	r.Redirect(req.URL.Query().Get(sessionauth.RedirectParam))
}

func SignupAccounts(r render.Render, html html.HTMLContext) {
	log.Debugln("Signup action !!!!!")
	r.HTML(200, "accounts/signup", html)
}

func SaveSignupAccounts(r render.Render, html html.HTMLContext, session sessions.Session, form models.User, errs binding.Errors, req *http.Request) {
	log.Debugln("SaveSignup action !!!!!")

	if len(errs) > 0 {
		log.Warning(errs)

		html["Errors"] = errs
		r.HTML(200, "accounts/signup", html)
		return
	}

	var errors binding.Errors

	if form.Password != form.Repassword {
		errors = append(errors, binding.Error{Message: "パスワードが一致していません"})
		html["Errors"] = errors
		r.HTML(200, "accounts/signup", html)
		return
	}

	user := models.User{}
	// err := models.DB.Where("email = ? AND password = ?", form.Email, form.Password).First(&user).Error
	// if err != nil {
	// log.Warning(err)

	// html["Errors"] = errors
	// r.HTML(200, "accounts/signup", html)
	// return
	// }

	err := sessionauth.AuthenticateSession(session, &user)
	if err != nil {
		log.Warning(err)

		html["Errors"] = errors
		r.HTML(200, "accounts/signup", html)
		return
	}

	r.Redirect(req.URL.Query().Get(sessionauth.RedirectParam))
}
