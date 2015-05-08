package configs

import (
	"github.com/k0kubun/pp"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type settings struct {
	Name          string
	Dsn           string
	Email         string
	AppName       string
	Keywords      string
	Copyright     string
	Description   string
	GroupServices []struct {
		Title string
		Href  string
		Src   string
	}
}

var Settings settings

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

	pp.Println(s)

	Settings = s
}
