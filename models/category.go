package models

import "database/sql"

type Category struct {
	Model

	Prefix string `sql:"type:varchar(8);index;not null"`

	Name string `sql:"type:varchar(255);not null"` // gin index
	Kana string `sql:"type:varchar(255);"`         // gin index

	Image   *Image
	ImageID sql.NullInt64

	Dicts []*Dict
}
