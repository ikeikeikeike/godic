package configs

import (
	"os"
	"path"

	"github.com/go-martini/martini"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type settings struct {
	Name          string
	Dsn           string
	Email         string
	Author        string
	AppName       string
	Keywords      string
	Copyright     string
	Description   string
	GroupServices []struct {
		Title string
		Href  string
		Src   string
	}

	RedisConn string
	UserAgent string
}

var Settings settings
var s = defaultSettings()

func defaultSettings() int {
	p, _ := os.Getwd()
	Register(
		path.Join(p, "config/settings.lua"),
		martini.Env,
	)
	return 0
}

func Register(filepath, environ string) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(filepath); err != nil {
		panic(err)
	}

	var s settings
	table := L.GetGlobal(environ).(*lua.LTable)
	if err := gluamapper.Map(table, &s); err != nil {
		panic(err)
	}

	Settings = s
}
