package models

import "database/sql"

type Category struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`         // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Prefix string `sql:"type:varchar(8);index;not null"`

	Image   *Image
	ImageID sql.NullInt64

	Dicts []*Dict
}
