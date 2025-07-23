package storage_test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisIntegration_SetGet_Success(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	err := client.Set(ctx, "integration-key", "integration-val", 0).Err()
	require.NoError(t, err)

	val, err := client.Get(ctx, "integration-key").Result()
	require.NoError(t, err)
	require.Equal(t, "integration-val", val)
}

func TestRedisIntegration_SetGet_NotFound(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	err := client.Set(ctx, "integration-key", "integration-val", 0).Err()
	require.NoError(t, err)

	val, err := client.Get(ctx, "other-key").Result()
	require.NoError(t, err)
	require.Equal(t, "", val)
}
