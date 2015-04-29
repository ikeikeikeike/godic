package funcmaps

import "html/template"

var HelperFuncs []template.FuncMap

func init() {
	HelperFuncs = []template.FuncMap{
		{
			"toAge":      ToAge,
			"safeHTML":   SafeHTML,
			"escapeHTML": EscapeHTML,
			"nl2br":      Nl2br,
			"toList":     List,
			"timeSince":  TimeSince,
		},
	}
}
