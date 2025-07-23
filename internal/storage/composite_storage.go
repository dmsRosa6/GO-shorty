package storage

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type CompositeStorage struct {
	RedisStorage    Storage
	InMemoryStorage Storage
}

func NewCompositeStorage(size int, addr string, expiration time.Duration) *CompositeStorage {
	return &CompositeStorage{
		RedisStorage:    NewRedisStorage(addr, expiration),
		InMemoryStorage: NewInMemoryStorage(size),
	}
}

func (cs *CompositeStorage) Set(ctx context.Context, key, value string) error {

	err := cs.RedisStorage.Set(ctx, key, value)

	if err != nil {
		return err
	}

	err = cs.InMemoryStorage.Set(ctx, key, value)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CompositeStorage) Get(ctx context.Context, key string) (string, error) {
	val, err := cs.InMemoryStorage.Get(ctx, key)

	if err != nil && err != ErrKeyNotFound {
		return "", err
	}

	if val != "" {
		return val, nil
	}

	val, err = cs.RedisStorage.Get(ctx, key)

	if err != nil && err != redis.Nil {
		return "", err
	}

	return val, nil
}
