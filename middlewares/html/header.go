package html

import "github.com/ikeikeikeike/godic/models/dict"

func HTMLHeader(html HTMLContext) {
	html["RecentDicts"] = dict.CachedDicts(10)
}
