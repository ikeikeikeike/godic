package redis

import (
	"fmt"

	"github.com/ikeikeikeike/godic/modules/configs"
	"github.com/ikeikeikeike/gopkg/redis"
)

var RC *redis.Client

func init() {
	conn := fmt.Sprintf(`{"conn": "%s"}`, configs.Settings.RedisConn)

	RC = redis.NewClient()
	err := RC.Initialize(conn)
	if err != nil {
		panic(err)
	}
}
