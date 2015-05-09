package html

import (
	"net/http"

	"github.com/go-martini/martini"
)

func RequestParams(html HTMLContext, p martini.Params, req *http.Request) {
	params := make(map[string]string)

	for k, v := range req.URL.Query() {
		params[k] = v[0]
	}
	for k, v := range p {
		params[k] = v
	}
	html["Params"] = params
}
