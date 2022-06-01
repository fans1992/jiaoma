// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"github.com/fans1992/jiaoma/pkg/cache"
	"github.com/fans1992/jiaoma/pkg/config"
)

// SetupCache 缓存
func SetupCache() {

	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
