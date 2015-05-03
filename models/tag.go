package models

import "database/sql"

type Tag struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`                // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Image   *Image
	ImageID sql.NullInt64

	Dicts []*Dict `gorm:"many2many:dict_tags;"`
}
