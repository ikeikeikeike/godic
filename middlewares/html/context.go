package html

import "github.com/go-martini/martini"

func GenHTMLContext() martini.Handler {
	var ctx HTMLContext = HTMLContext{}
	return func(c martini.Context) {
		c.Map(ctx)
	}
}
