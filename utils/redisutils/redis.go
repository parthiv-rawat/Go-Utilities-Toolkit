package redisutils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisSet(key, val string, ttl time.Duration) error {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer client.Close()
	return client.Set(context.Background(), key, val, ttl).Err()
}

func RedisGet(key string) (string, error) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer client.Close()
	return client.Get(context.Background(), key).Result()
}
