package models

import (
	"database/sql"
)

type Dict struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`                // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Content string `sql:"type:text"` // gin index

	Prefix string `sql:"type:varchar(8);index;not null"`

	Image   *Image
	ImageID sql.NullInt64

	Category   *Category
	CategoryID sql.NullInt64

	Tags []*Tag `gorm:"many2many:dict_tags;"`
}

func (m *Dict) BeforeCreate() error {
	m.Prefix = letterCombinePtn(7)
	return nil
}

func (m *Dict) GetPrefix() string {
	if m.Category != nil {
		return "./" + m.Category.Prefix + "/" + m.Prefix
	} else {
		return "./" + m.Prefix
	}
}
