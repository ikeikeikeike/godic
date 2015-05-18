package views

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/models/category"
	"github.com/ikeikeikeike/godic/models/dict"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func Roots(r render.Render, s sessions.Session, html html.HTMLContext) {
	log.Debugln("Roots action !!!!!")

	var modified []*models.Dict
	dict.RelationDB().Limit(10).Order("dicts.updated_at DESC").Find(&modified)
	html["ModifiedDicts"] = latests

	var latests []*models.Dict
	dict.RelationDB().Limit(10).Order("dicts.created_at DESC").Find(&latests)
	html["LatestDicts"] = latests

	html["Categories"] = category.CategoriesALL()

	r.HTML(200, "roots/index", html)
}

func LatestRoots(r render.Render, html html.HTMLContext, req *http.Request) {
	log.Debugln("LatestRoots action !!!!!")

	pers := 25
	db := dict.RelationDB()

	var max int
	db.Count(&max)

	pager := pagination.NewPaginator(req, pers, max)
	html["Paginator"] = &pager

	var latests []*models.Dict
	db.Limit(pers).Offset(pager.Offset()).Order("dicts.id DESC").Find(&latests)
	for _, d := range latests {
		d.TagsLoader()
	}

	html["LatestDicts"] = latests

	r.HTML(200, "roots/latest", html)
}

func ModifiedRoots(r render.Render, html html.HTMLContext, req *http.Request) {
	log.Debugln("ModifiedRoots action !!!!!")

	pers := 25
	db := dict.RelationDB()

	var max int
	db.Count(&max)

	pager := pagination.NewPaginator(req, pers, max)
	html["Paginator"] = &pager

	var modified []*models.Dict
	db.Limit(pers).Offset(pager.Offset()).Order("dicts.updated_at DESC").Find(&modified)
	for _, d := range modified {
		d.TagsLoader()
	}

	html["ModifiedDicts"] = modified

	r.HTML(200, "roots/modified", html)
}
