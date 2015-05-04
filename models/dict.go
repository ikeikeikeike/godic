package models

import (
	"crypto/rand"
	"database/sql"
)

type Dict struct {
	Model

	Name   string `sql:"type:varchar(255);unique;not null"` // gin index
	Yomi   string `sql:"type:varchar(255);"`                // gin index
	Romaji string `sql:"type:varchar(128)"`
	Gyou   string `sql:"type:varchar(6);index"`

	Outline string `sql:"type:text"` // gin index

	Prefix string `sql:"type:varchar(8);index;not null"`

	Image   *Image
	ImageID sql.NullInt64

	Category   *Category
	CategoryID sql.NullInt64

	Tags []*Tag `gorm:"many2many:dict_tags;"` // Many-To-Many relationship, 'user_languages' is join table
}

func (m *Dict) BeforeCreate() error {
	m.Prefix = randPrefix(7)
	return nil
}

func (m *Dict) GetPrefix() string {
	if m.Category != nil {
		return "./" + m.Category.Prefix + "/" + m.Prefix
	} else {
		return "./" + m.Prefix
	}
}

func randPrefix(n int) string {
	const letters = "abcdefg" // 7P7=7*6*5*4*3*2*1=5040
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
