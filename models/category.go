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
	ImageID sql.NullInt64 `sql:"index"`

	Dicts []*Dict
}

func (m *Category) LatestDicts(limit int) {
	DB.Model(&m).Preload("Image").Preload("Category").Preload("Comments").
		Order("dicts.id DESC").Limit(limit).Related(&m.Dicts, "Dicts")
}

func (m *Category) ModifiedDicts(limit int) {
	DB.Model(&m).Preload("Image").Preload("Category").Preload("Comments").
		Order("dicts.updated_at DESC").Limit(limit).Related(&m.Dicts, "Dicts")
}
