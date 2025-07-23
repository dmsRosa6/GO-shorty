package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client     *redis.Client
	expiration time.Duration
}

func NewRedisStorage(addr string, expiration time.Duration) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStorage{client: client, expiration: expiration}
}

func (redis *RedisStorage) Set(ctx context.Context, key, value string) error {
	return redis.client.Set(ctx, key, value)
}


func (redis *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	return redis.client.Get(ctx, key).Result()
}