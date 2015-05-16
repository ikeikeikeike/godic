package views

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	// "github.com/martini-contrib/oauth2"
	// goauth2 "golang.org/x/oauth2"

	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/modules/funcmaps"
	"github.com/ikeikeikeike/godic/views/forms"
)

var App *martini.ClassicMartini

func init() {
	App = martini.Classic()

	funcs := append(funcmaps.HelperFuncs, template.FuncMap{"urlFor": App.URLFor})
	App.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
		Funcs:      funcs,
	}))

	App.Use(html.GenHTMLContext())
	App.Use(html.HTMLHeader)
	App.Use(html.HTMLMeta)
	App.Use(html.HTMLSettings)

	App.Use(sessions.Sessions("godic_sesssion", sessions.NewCookieStore([]byte("secret09131ffl2"))))
	App.Use(sessionauth.SessionUser(models.GenerateAnonymousUser))

	// App.Use(oauth2.Github(&goauth2.Config{
	// ClientID:     "0.0",
	// ClientSecret: "o.o",
	// // RedirectURL:  "http://localhost:3000/oauth2callback",
	// Scopes: []string{"user:email", "read:org"},
	// }))
	// App.Use(oauth2.Facebook(&goauth2.Config{
	// ClientID:     "0.0",
	// ClientSecret: "o.o",
	// // RedirectURL:  "http://localhost:3000/oauth2callback",
	// Scopes: []string{"user:email", "read:org"},
	// }))
	// App.Use(oauth2.Google(&goauth2.Config{
	// ClientID:     "",
	// ClientSecret: "",
	// // RedirectURL:  "http://localhost:3000/oauth2callback",
	// Scopes: []string{"https://www.googleapis.com/auth/drive"},
	// }))

	allowCORS := cors.Allow(&cors.Options{
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	})

	App.Group("", func(r martini.Router) {
		r.Get("/", Roots).Name("roots")
		r.Get("/latest", LatestRoots).Name("roots_latest")
		r.Get("/modified", ModifiedRoots).Name("roots_modified")
	})

	App.Group("", func(r martini.Router) {
		r.Get("/signup", SignupAccounts).Name("accounts_signup")
		r.Get("/login", LoginAccounts).Name("accounts_login")
		r.Post("/login", binding.Form(models.User{}), SaveLoginAccounts).Name("accounts_login")
		r.Get("/logout", sessionauth.LoginRequired, func(r render.Render, session sessions.Session, user sessionauth.User) {
			sessionauth.Logout(session, user)
			r.Redirect("/")
		})
	})

	App.Group("/abouts", func(r martini.Router) {
		r.Get("/sitemap", func(r render.Render) { r.Redirect("/") }).Name("abouts_sitemap")
	})

	App.Group("/rss", func(r martini.Router) {
		r.Get("/latest.xml", LatestRSS).Name("rss_latest")
		r.Get("/modified.xml", ModifiedRSS).Name("rss_modified")
	}, allowCORS)

	App.Group("/category", func(r martini.Router) {
		r.Get("/:name", Categories).Name("categories")
	})

	App.Group("/d", func(r martini.Router) {
		r.Get("/index", func(r render.Render) { r.Redirect("/") }).Name("index")
		r.Get("/new/", sessionauth.LoginRequired, NewDicts).Name("new")
		r.Get("/new/:name", sessionauth.LoginRequired, NewDicts).Name("new")
		r.Get("/:name", ShowDicts).Name("show")
		r.Get("/edit/:name", sessionauth.LoginRequired, EditDicts).Name("edit")
		r.Get("/history/:name", DictsHistory).Name("history")
		r.Get("/:name/:sha1", ShowDicts).Name("show")
		r.Get(`/compare/:name/(?P<fromsha1>[^\.]+)\.{2,3}(?P<tosha1>.+)`, CompareDicts).Name("compare")
	}, html.RequestParams)

	App.Group("/_d", func(r martini.Router) {
		r.Get("/:name", func(r render.Render, p martini.Params) { r.Redirect("/d/" + p["name"]) })
		r.Put("/:name", binding.Bind(forms.Commit{}), UpdateDicts).Name("api_put")
		r.Post("/:name", binding.Bind(forms.Commit{}), CreateDicts).Name("api_post")
		r.Delete("/:name", DeleteDicts).Name("api_delete")
	})

	App.NotFound(func(r render.Render, html html.HTMLContext) {
		r.HTML(404, "errors/404", html)
	})
}
