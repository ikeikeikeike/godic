package html

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ikeikeikeike/godic/modules/configs"
	"github.com/ikeikeikeike/gopkg/convert"
	woothee "github.com/woothee/woothee-go"
)

type Meta struct {
	AppName     string
	Email       string
	Copyright   string
	Keywords    string
	Description string
	URL         string
	Host        string
	Domain      string
	UA          *woothee.Result
}

func NewMeta() *Meta {
	return &Meta{
		AppName:     "",
		Email:       "",
		Copyright:   "",
		Keywords:    "",
		Description: "",
		URL:         "",
		Host:        "",
		Domain:      "",
	}
}

func HTMLMeta(res http.ResponseWriter, req *http.Request, html HTMLContext) {
	m := NewMeta()

	m.AppName = configs.Settings.AppName
	m.Email = configs.Settings.Email
	m.Copyright = configs.Settings.Copyright
	m.Keywords = configs.Settings.Keywords
	m.Description = configs.Settings.Description

	m.URL = BuildRequestUrl(req, "")
	m.Host = Site(req)
	m.Domain = Domain(req)
	m.UA, _ = woothee.Parse(UserAgent(req))

	html["Meta"] = m
}

func BuildRequestUrl(req *http.Request, uri string) string {
	if uri == "" {
		uri = req.RequestURI
	}
	return fmt.Sprintf("%s:%s%s", Site(req), convert.ToStr(Port(req)), uri)
}

func Port(req *http.Request) int {
	parts := strings.Split(req.Host, ":")
	if len(parts) == 2 {
		port, _ := strconv.Atoi(parts[1])
		return port
	}
	return 80
}

func Site(req *http.Request) string {
	return Scheme(req) + "://" + Domain(req)
}

func Scheme(req *http.Request) string {
	if req.URL.Scheme != "" {
		return req.URL.Scheme
	}
	if req.TLS == nil {
		return "http"
	}
	return "https"
}

func Domain(req *http.Request) string {
	return Host(req)
}

func Host(req *http.Request) string {
	if req.Host != "" {
		hostParts := strings.Split(req.Host, ":")
		if len(hostParts) > 0 {
			return hostParts[0]
		}
		return req.Host
	}
	return "localhost"
}

func UserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}
