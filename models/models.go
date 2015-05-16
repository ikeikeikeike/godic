package models

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/modules/configs"
	"github.com/ikeikeikeike/gopkg/rdm"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time `sql:"index"`
	UpdatedAt time.Time `sql:"index"`
	DeletedAt *time.Time `sql:"index"`
}

var DB gorm.DB

func init() {
	var err error

	dialect := "sqlite3"
	if strings.HasPrefix(configs.Settings.Dsn, "postgres") {
		dialect = "postgres"
	}

	switch martini.Env {
	case "production":
		DB, err = gorm.Open(dialect, configs.Settings.Dsn)
		if err != nil {
			panic(err)
		}

		DB.DB()
		DB.DB().Ping()
		DB.DB().SetMaxIdleConns(100)
		DB.DB().SetMaxOpenConns(100)
	default:
		DB, err = gorm.Open(dialect, configs.Settings.Dsn)
		if err != nil {
			panic(err)
		}

		DB.DB()
		DB.LogMode(true)
	}

	DB.AutoMigrate(&Dict{}, &Tag{}, &Category{}, &Image{}, &User{}, &Comment{})
	DB.Table("dict_tags").AddIndex("idx_dict_tags_dict_id", "dict_id")
	DB.Table("dict_tags").AddIndex("idx_dict_tags_tag_id", "tag_id")

	InitSeed()
}

func InitSeed() {
	cate := new(Category)
	DB.Where(Category{Name: "ノンカテゴリー"}).
		Attrs(Category{Yomi: "のんかてごり", Romaji: "nonkategori", Gyou: "no", Prefix: "none"}).
		FirstOrCreate(&cate)

	cate = new(Category)
	DB.Where(Category{Name: "アイドル・女優"}).
		Attrs(Category{Yomi: "あいどるじょゆう", Romaji: "aidorujoyu", Gyou: "a", Prefix: "diva"}).
		FirstOrCreate(&cate)

	cate = new(Category)
	DB.Where(Category{Name: "漫画・アニメ"}).
		Attrs(Category{Yomi: "まんがあにめ", Romaji: "mangaanime", Gyou: "ma", Prefix: "anime"}).
		FirstOrCreate(&cate)

	cate = new(Category)
	DB.Where(Category{Name: "漫画・アニメキャラ"}).
		Attrs(Category{Yomi: "まんがあにめきゃら", Romaji: "mangaanimekyara", Gyou: "ma", Prefix: "character"}).
		FirstOrCreate(&cate)

	if martini.Env != "production" {

		basePath, _ := os.Getwd()
		bytes, err := ioutil.ReadFile(path.Join(basePath, "template.txt"))

		i := 0
		for i < 0 {
			name := letterCombinePtn(5)

			content := ""
			if err == nil {
				content = fmt.Sprintf(string(bytes), name, "")
			}

			c := &Category{}
			DB.First(c, int64(rdm.RandomNumber(1, 5)))

			d := new(Dict)
			DB.Where(Dict{Name: name}).
				Attrs(Dict{
				Yomi:     letterCombinePtn(3),
				Prefix:   letterCombinePtn(7),
				Content:  content,
				Category: c,
			}).
				FirstOrCreate(&d)
			i++
		}
	}
}

func letterCombinePtn(n int) string {
	const letters = "abcdefg" // 7P7=7*6*5*4*3*2*1=5040
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
