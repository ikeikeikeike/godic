package html

import (
	"net/http"

	"github.com/ikeikeikeike/godic/modules/configs"
)

func HTMLSettings(res http.ResponseWriter, req *http.Request, html HTMLContext) {
	html["Settings"] = configs.Settings
}
