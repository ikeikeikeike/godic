package models

import (
	"database/sql"
	"encoding/json"
	"reflect"

	"github.com/ikeikeikeike/godic/modules/funcmaps"
	"github.com/ikeikeikeike/godic/modules/redis"
)

type Dict struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`                // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Content     string `sql:"type:text"` // gin index
	ContentHTML string `sql:"type:text"` // gin index

	Prefix string `sql:"type:varchar(16);index;not null"`

	Image   *Image
	ImageID sql.NullInt64 `sql:"index"`

	Category   *Category
	CategoryID sql.NullInt64 `sql:"index"`

	Tags []*Tag `gorm:"many2many:dict_tags"`
}

func (m *Dict) BeforeCreate() error {
	m.Prefix = letterCombinePtn(7)
	return nil
}

func (m *Dict) BeforeSave() error {
	html := funcmaps.MarkdownHTML(m.Content)

	do := false
	for _, img := range funcmaps.ExtractIMGs(html) {
		if m.Image == nil {
			m.Image = NewImageByIMG(img)
			do = true
		} else if m.Image.Width < img.Width || m.Image.Height < img.Height {
			m.Image = NewImageByIMG(img)
			do = true
		}
	}
	if do {
		DB.Save(m.Image)
		m.ImageID = sql.NullInt64{m.Image.ID, true}
	}

	c := funcmaps.AutoLink(html, cachedDictNames())
	m.ContentHTML = c

	return nil
}

func (m *Dict) GetPrefix() string {
	if m.Category != nil {
		return "./" + m.Category.Prefix + "/" + m.Prefix
	} else {
		return "./" + m.Prefix
	}
}

func cachedDictNames() []string {
	key := "godic.models.dict.caches.CachedDicts"
	s := reflect.ValueOf(redis.RC.Get(key))

	var dicts []*Dict

	if !redis.RC.IsExist(key) {
		DB.Table("dicts").Limit(-1).Find(&dicts)

		bytes, _ := json.Marshal(dicts)
		redis.RC.Put(key, bytes, 60*60*24*1)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &dicts)
	}

	names := make([]string, len(dicts))
	for i, d := range dicts {
		names[i] = d.Name
	}

	return names
}
