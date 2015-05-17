package views

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/middlewares/html"
	"github.com/ikeikeikeike/godic/models"
	"github.com/martini-contrib/render"
)

func Categories(r render.Render, params martini.Params, html html.HTMLContext, req *http.Request) {
	log.Debugln("Categories action !!!!!")

	var list []*models.Category
	err := models.DB.Where("categories.prefix = ?", params["name"]).Find(&list).Error
	if err != nil || len(list) <= 0 {
		r.HTML(404, "errors/404", html)
		return
	}

	pers := 25
	db := models.DB.Table("dicts").
		Preload("Image").Preload("Category").Preload("Comments").
		Select("dicts.*").
		Joins("INNER JOIN categories ON categories.id = dicts.category_id").
		Where("categories.prefix = ?", params["name"])

	var max int
	db.Count(&max)

	pager := pagination.NewPaginator(req, pers, max)
	html["Paginator"] = &pager

	var dicts []*models.Dict
	db.Limit(pers).Offset(pager.Offset()).Order("dicts.updated_at DESC").Find(&dicts)
	for _, d := range dicts {
		d.TagsLoader()
	}

	html["Dicts"] = dicts

	r.HTML(200, "categories/index", html)
}
