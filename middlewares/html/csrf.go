package html

import "github.com/martini-contrib/csrf"

func HTMLCSRF(x csrf.CSRF, html HTMLContext) {
	html["CSRF"] = x.GetToken()
}
