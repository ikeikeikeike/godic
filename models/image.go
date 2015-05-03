package models

type Image struct {
	Model

	Name string `sql:"type:varchar(255);not null"`
	Src  string `sql:"type:varchar(255);not null"`

	Ext    string
	Mime   string
	Width  int
	Height int
}
