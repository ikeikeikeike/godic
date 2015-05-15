package dict

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ikeikeikeike/godic/models"
	"github.com/ikeikeikeike/godic/modules/redis"
)

/*
	Cache expires: 1 hour
*/
func CachedDicts(limit int) (objs []*models.Dict) {
	key := fmt.Sprintf("godic.models.dict.caches.CachedDicts:limit%d", limit)
	s := reflect.ValueOf(redis.RC.Get(key))

	if !redis.RC.IsExist(key) {
		RelationDB().Limit(limit).Order("dicts.updated_at DESC").Find(&objs)

		bytes, _ := json.Marshal(objs)
		redis.RC.Put(key, bytes, 60*60*1)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &objs)
	}

	return
}

func CachedNames() []string {
	dicts := CachedDicts(-1)
	names := make([]string, len(dicts))
	for i, d := range dicts {
		names[i] = d.Name
	}

	return names
}
