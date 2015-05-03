package models

import (
	"os"
	"time"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

var DB gorm.DB

func init() {
	var err error

	switch martini.Env {
	case "production":
		DB, err = gorm.Open("postgres", os.Getenv("DSN"))
		if err != nil {
			panic(err)
		}

		DB.DB()
		DB.DB().Ping()
		DB.DB().SetMaxIdleConns(100)
		DB.DB().SetMaxOpenConns(100)
	default:
		DB, err = gorm.Open("sqlite3", os.Getenv("DSN"))
		if err != nil {
			panic(err)
		}

		DB.DB()
		DB.LogMode(true)
	}

	DB.AutoMigrate(&Dict{}, &Tag{}, &Category{}, &Image{})
}
