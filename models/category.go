package models

import "database/sql"

type Category struct {
	Model

	Name   string `sql:"type:varchar(255);not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`         // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Prefix string `sql:"type:varchar(16);unique;not null"`

	Image   *Image
	ImageID sql.NullInt64

	Dicts []*Dict
}

func (m *Category) DictLoader(limit int) {
	DB.Model(&m).Preload("Category").Preload("Image").
		Order("dicts.id DESC").Limit(limit).Related(&m.Dicts, "Dicts")
}
