package storage_test

import (
	"context"
	"testing"

	"github.com/dmsRosa6/go-shorty/internal/storage"
	"github.com/stretchr/testify/require"
)

func Test_InMemorySet_Success(t *testing.T) {
	mem := storage.NewInMemoryStorage(1)

	err := mem.Set(context.Background(), "key", "value")
	require.NoError(t, err)
}

func Test_InMemoryget_Sucess(t *testing.T) {
	mem := storage.NewInMemoryStorage(1)

	err := mem.Set(context.Background(), "key", "value")
	require.NoError(t, err)
	get, err := mem.Get(context.Background(), "key")
	require.NoError(t, err)

	require.Equal(t, get, "value")
}

func Test_InMemoryget_NotFound(t *testing.T) {
	mem := storage.NewInMemoryStorage(1)

	err := mem.Set(context.Background(), "key", "value")
	require.NoError(t, err)

	_, err = mem.Get(context.Background(), "otherkey")
	require.Error(t, err, "key not found")

}

func Test_InMemorySet_Success_Override(t *testing.T) {
	mem := storage.NewInMemoryStorage(1)

	err := mem.Set(context.Background(), "key", "first")
	require.NoError(t, err)
	err = mem.Set(context.Background(), "key", "second")
	require.NoError(t, err)

	get, err := mem.Get(context.Background(), "key")
	require.NoError(t, err)

	require.Equal(t, get, "second")
}

func Test_InMemorySet_Success_Eviction(t *testing.T) {
	mem := storage.NewInMemoryStorage(1)

	err := mem.Set(context.Background(), "first", "first")
	require.NoError(t, err)
	err = mem.Set(context.Background(), "second", "second")
	require.NoError(t, err)

	_, err = mem.Get(context.Background(), "first")
	require.Error(t, err, "key not found")

	get, err := mem.Get(context.Background(), "second")
	require.NoError(t, err)

	require.Equal(t, get, "second")
}

func Test_InMemorySet_EvictionOrder(t *testing.T) {
	mem := storage.NewInMemoryStorage(2)

	require.NoError(t, mem.Set(context.Background(), "a", "A"))
	require.NoError(t, mem.Set(context.Background(), "b", "B"))
	_, err := mem.Get(context.Background(), "a")
	require.NoError(t, err)

	require.NoError(t, mem.Set(context.Background(), "c", "C"))

	_, err = mem.Get(context.Background(), "b")
	require.Error(t, err, "key not found")

	val, err := mem.Get(context.Background(), "a")
	require.NoError(t, err)
	require.Equal(t, "A", val)

	val, err = mem.Get(context.Background(), "c")
	require.NoError(t, err)
	require.Equal(t, "C", val)
}

func Test_InMemorySet_CapacityLimit(t *testing.T) {
	mem := storage.NewInMemoryStorage(2)

	require.NoError(t, mem.Set(context.Background(), "a", "A"))
	require.NoError(t, mem.Set(context.Background(), "b", "B"))
	require.NoError(t, mem.Set(context.Background(), "c", "C"))

	_, err := mem.Get(context.Background(), "a")
	_, err2 := mem.Get(context.Background(), "b")
	_, err3 := mem.Get(context.Background(), "c")

	count := 0
	if err == nil {
		count++
	}
	if err2 == nil {
		count++
	}
	if err3 == nil {
		count++
	}
	require.Equal(t, 2, count)
}
