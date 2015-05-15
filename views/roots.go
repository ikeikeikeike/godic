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
)

func Roots(r render.Render, html html.HTMLContext) {
	log.Debugln("Roots action !!!!!")

	latests := category.CategoriesALL()
	for _, c := range latests {
		c.LatestDicts(10)
	}
	html["LatestDicts"] = latests

	modified := category.CategoriesALL()
	for _, c := range modified {
		c.ModifiedDicts(10)
	}
	html["ModifiedDicts"] = modified

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
