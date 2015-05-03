package models

import "database/sql"

type Dict struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`                // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Outline string `sql:"type:text"` // gin index

	Image   *Image
	ImageID sql.NullInt64

	Category   *Category
	CategoryID sql.NullInt64

	Tags []*Tag `gorm:"many2many:dict_tags;"` // Many-To-Many relationship, 'user_languages' is join table
}

func (m *Dict) GetPrefix() string {
	if m.Category != nil {
		return "./" + m.Category.Prefix
	} else {
		return "./"
	}
}
