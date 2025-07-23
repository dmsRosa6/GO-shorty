package storage_test

import (
	"context"
	"errors"
	"testing"

	redismock "github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisSet_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectSet("key", "value", 0).SetVal("OK")

	err := db.Set(context.Background(), "key", "value", 0).Err()
	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisSet_InternalError(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectSet("key", "value", 0).SetErr(errors.New("redis error"))

	_, err := db.Set(context.Background(), "key", "value", 0).Result()
	require.Error(t, err)

	require.EqualError(t, err, "redis error")
}

func TestRedisGet_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("key").SetVal("value")

	val, err := db.Get(context.Background(), "key").Result()
	require.NoError(t, err)

	require.Equal(t, val, "value")
}

func TestRedisGet_NotFound(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("key").SetVal("")

	val, err := db.Get(context.Background(), "key").Result()
	require.NoError(t, err)

	require.NotEqual(t, val, "some value")
}

func TestRedisGet_InternalError(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("key").SetErr(errors.New("redis error"))

	_, err := db.Get(context.Background(), "key").Result()
	require.Error(t, err)

	require.EqualError(t, err, "redis error")
}

