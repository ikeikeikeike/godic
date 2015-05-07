package category

import (
	m "github.com/ikeikeikeike/godic/models"
	"github.com/jinzhu/gorm"
)

// Scope

func SetNone(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "none")
}

func SetDiva(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "diva")
}

func SetAnime(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "anime")
}

func SetCharacter(db *gorm.DB) *gorm.DB {
	return db.Where("categories.prefix = ?", "character")
}

// get alls

func RelationDB() *gorm.DB {
	return m.DB.Table("categories").
		// Preload("Dicts").  XXX: See category.go#DictLoader
		Select("categories.*")
}

func CategoriesALL() (list []*m.Category) {
	RelationDB().Order("categories.id DESC").Find(&list)
	return
}

func Categories() (list []*m.Category) {
	prefixes := []string{"diva", "anime", "character"}
	RelationDB().Where("categories.prefix in (?)", prefixes).Order("categories.id DESC").Find(&list)
	return
}

// get ones

func None() *m.Category {
	c := &m.Category{}
	RelationDB().Scopes(SetNone).First(c)
	return c
}

func Diva() *m.Category {
	c := &m.Category{}
	RelationDB().Scopes(SetDiva).First(c)
	return c
}

func Anime() *m.Category {
	c := &m.Category{}
	RelationDB().Scopes(SetAnime).First(c)
	return c
}

func Character() *m.Category {
	c := &m.Category{}
	RelationDB().Scopes(SetCharacter).First(c)
	return c
}
