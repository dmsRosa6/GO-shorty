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

func (redis *RedisStorage) Set(ctx context.Context, key string, value interface{}) error {
	//TODO:check what that string may or may not return
	_, err := redis.client.Set(ctx, key, value, redis.expiration).Result()

	return err
}

func (redis *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	return redis.client.Get(ctx, key).Result()
}
