package dict

import (
	"encoding/json"
	"reflect"

	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/modules/redis"
)

/*
	Cache expires: 1 days
*/
func CachedDicts() (objs []*models.Dict) {
	key := "godic.models.dict.caches.CachedDicts"
	s := reflect.ValueOf(redis.RC.Get(key))

	if !redis.RC.IsExist(key) {
		RelationDB().Limit(-1).Find(&objs)

		bytes, _ := json.Marshal(objs)
		redis.RC.Put(key, bytes, 60*60*24*1)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &objs)
	}

	return
}

func CachedNames() []string {
	dicts := CachedDicts()
	names := make([]string, len(dicts))
	for i, d := range dicts {
		names[i] = d.Name
	}

	return names
}
