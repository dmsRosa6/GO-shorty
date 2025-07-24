package storage

import (
	"context"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type CompositeStorage struct {
	RedisStorage    Storage
	InMemoryStorage Storage
}

func NewCompositeStorage(size int, addr string, expiration time.Duration) *CompositeStorage {
	redis, err := NewRedisStorage(addr, expiration)

	if err != nil {
		log.Printf("warning: failed to connect to Redis: %v", err)
	}

	return &CompositeStorage{
		RedisStorage:    redis,
		InMemoryStorage: NewInMemoryStorage(size),
	}
}

func (cs *CompositeStorage) Set(ctx context.Context, key, value string) error {

	if cs.RedisStorage != nil {
		err := cs.RedisStorage.Set(ctx, key, value)

		if err != nil {
			return err
		}
	}

	err := cs.InMemoryStorage.Set(ctx, key, value)

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

	if cs.RedisStorage != nil {
		val, err = cs.RedisStorage.Get(ctx, key)

		if err != nil && err != redis.Nil {
			return "", err
		}
	}

	return val, nil
}
