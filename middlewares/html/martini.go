package html

import "github.com/go-martini/martini"

func RequestParams(html HTMLContext, params martini.Params) {
	html["Params"] = params
}
