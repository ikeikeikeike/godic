package html

import (
	"net/http"

	dichttp "github.com/ikeikeikeike/godic/modules/http"
	woothee "github.com/woothee/woothee-go"
)

type Meta struct {
	URL    string
	Host   string
	Domain string
	UA     *woothee.Result
}

func NewMeta() *Meta {
	return &Meta{
		URL:    "",
		Host:   "",
		Domain: "",
	}
}

func HTMLMeta(res http.ResponseWriter, req *http.Request, html HTMLContext) {
	m := NewMeta()

	m.URL = dichttp.BuildRequestUrl(req, "")
	m.Host = dichttp.Site(req)
	m.Domain = dichttp.Domain(req)
	m.UA, _ = woothee.Parse(dichttp.UserAgent(req))

	html["Meta"] = m
}
