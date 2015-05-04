package category

import (
	m "github.com/ikeikeikeike/godic/models"
	"github.com/jinzhu/gorm"
)

// Scope

func SetDiva(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "diva")
}

func SetAnime(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "anime")
}

func SetCharacter(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "character")
}

func Categories() *gorm.DB {
	return m.DB.Table("categories").Preload("Dicts").Select("categories.*")
}

func Diva() *m.Category {
	c := &m.Category{}
	Categories().Scopes(SetDiva).First(c)
	return c
}

func Anime() *m.Category {
	c := &m.Category{}
	Categories().Scopes(SetAnime).First(c)
	return c
}

func Character() *m.Category {
	c := &m.Category{}
	Categories().Scopes(SetCharacter).First(c)
	return c
}
