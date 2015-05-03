package html

import (
	"github.com/go-martini/martini"
)

func GenHTMLContext() martini.Handler {
	return func(c martini.Context) {
		c.Map(HTMLContext{})
	}
}
