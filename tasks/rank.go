package tasks

import (
	"go-crud/cache"
)

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {

	return cache.RedisClient.Del("rank:daily").Err()
}
