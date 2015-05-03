package models

import "database/sql"

type Dict struct {
	Model

	Name string `sql:"type:varchar(255);unique;not null"` // gin index
	Kana string `sql:"type:varchar(255);"`                // gin index

	Outline string `sql:"type:text"` // gin index

	Image   *Image
	ImageID sql.NullInt64

	Category   *Category
	CategoryID sql.NullInt64
}

func (m *Dict) GetPrefix() string {
	if m.Category != nil {
		return "./" + m.Category.Prefix
	} else {
		return "./"
	}
}
