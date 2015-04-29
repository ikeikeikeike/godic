package html

import "github.com/go-martini/martini"

func SetParams(html HTMLContext, params martini.Params) {
	html["Params"] = params
}
