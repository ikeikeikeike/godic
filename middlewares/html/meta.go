package html

import (
	woothee "github.com/woothee/woothee-go"
)

type Meta struct {
	AppName         string
	Copyright       string
	Author          string
	Email           string
	Keywords        string
	Description     string
	ApplicationName string
	Domain          string
	Host            string
	Url             string
	Type            string
	Title           string
	Image           string
	SiteName        string
	Locale          string
	FBAppId         string
	TWCard          string
	TWDomain        string
	TWSite          string
	TWImage         string
	UA              *woothee.Result
}

func NewMeta() *Meta {
	return &Meta{
		AppName:     "",
		Copyright:   "",
		Author:      "",
		Email:       "",
		Keywords:    "",
		Description: "",
		Domain:      "",
		Host:        "",
		Url:         "",
		Type:        "",
		Title:       "",
		Image:       "",
		SiteName:    "",
		Locale:      "",
		FBAppId:     "",
		TWCard:      "",
		TWDomain:    "",
		TWSite:      "",
		TWImage:     "",
	}
}

func HTMLMeta(html HTMLContext) {
	m := NewMeta()
	// m.Url = c.BuildRequestUrl("")
	// m.Host = c.Ctx.Input.Site()
	// m.UA, _ = woothee.Parse(c.Ctx.Input.UserAgent())
	html["Meta"] = m
}
