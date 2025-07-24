package storage

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var failedConnection = errors.New("failed to connect to Redis")

type RedisStorage struct {
	client     *redis.Client
	expiration time.Duration
}

func NewRedisStorage(addr string, expiration time.Duration) (Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Attempt to ping Redis
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, failedConnection
	}

	return &RedisStorage{client: client, expiration: expiration}, nil
}

func (redis *RedisStorage) Set(ctx context.Context, key string, value string) error {
	//TODO:check what that string may or may not return
	_, err := redis.client.Set(ctx, key, value, redis.expiration).Result()

	return err
}

func (redis *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	return redis.client.Get(ctx, key).Result()
}
