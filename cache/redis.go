package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

func init() {
	initRedisClient(RedisCacheConfig{
		Host: "localhost:6379",
		Db: 0,
		Exp: 1,
	})
}

var redisClient *redis.Client

type RedisCacheConfig struct {
	Host string
	Db int
	Exp time.Duration
}

func initRedisClient(config RedisCacheConfig) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.Host,
		DB: config.Db,
		Password: "",
	})
}

func GetRedisClient() (*redis.Client) {
	return redisClient
}