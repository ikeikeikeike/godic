package models

import "github.com/ikeikeikeike/godic/modules/funcmaps"

type Image struct {
	Model

	Name string `sql:"type:varchar(255);not null"`
	Src  string `sql:"type:varchar(255);not null"`

	Ext    string
	Mime   string
	Width  int
	Height int
}

func NewImageByIMG(img *funcmaps.Img) *Image {
	return &Image{
		Name:   img.Alt,
		Src:    img.Src,
		Ext:    img.Ext,
		Mime:   img.Mime,
		Width:  img.Width,
		Height: img.Height,
	}
}
