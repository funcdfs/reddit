package redis

import (
	"fmt"
	"reddit/settings"

	"github.com/go-redis/redis"
)

// Declare a global rdb variable
var rdb *redis.Client

// Init init connection use viper config
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // use password or "".
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
