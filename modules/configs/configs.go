package configs

import (
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type settings struct {
	Name        string
	Dsn         string
	Email       string
	AppName     string
	Keywords    string
	Copyright   string
	Description string
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

	Settings = s
}
