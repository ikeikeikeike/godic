package models

import "database/sql"

type Category struct {
	Model

	Prefix string `sql:"type:varchar(8);index;not null"`

	Name   string `sql:"type:varchar(255);not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`         // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Image   *Image
	ImageID sql.NullInt64

	Dicts []*Dict
}
