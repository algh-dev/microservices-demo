package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func init() {
	InitRedisClient(RedisCacheConfig{
		Host: "localhost:6379",
		Db: 0,
		Exp: 1,
	})
}

var redisClient *redis.Client

var ctx = context.Background()

type RedisCacheConfig struct {
	Host string
	Db int
	Exp time.Duration
}

func InitRedisClient(config RedisCacheConfig) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.Host,
		DB: config.Db,
		Password: "",
	})
}

func GetRedisClient() (*redis.Client, context.Context) {
	return redisClient, ctx
}