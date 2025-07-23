package storage

import (
	"container/list"
	"context"
	"errors"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

type MapEntry struct {
	key   string
	value string
}

type InMemoryStorage struct {
	mu      sync.RWMutex
	store   map[string]*list.Element
	lru     *list.List
	maxSize int
}

func NewInMemoryStorage(size int) Storage {
	return &InMemoryStorage{
		store:   make(map[string]*list.Element, size),
		maxSize: size,
		lru:     list.New(),
	}
}

func (mem *InMemoryStorage) Set(ctx context.Context, key, value string) error {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	if elem, ok := mem.store[key]; ok {
		elem.Value = MapEntry{key: key, value: value}
		mem.lru.MoveToFront(elem)
		return nil
	}

	if len(mem.store) >= mem.maxSize {
		back := mem.lru.Back()
		if back != nil {
			mem.lru.Remove(back)
			delete(mem.store, back.Value.(MapEntry).key)
		}
	}

	elem := mem.lru.PushFront(MapEntry{key: key, value: value})
	mem.store[key] = elem
	return nil
}

func (mem *InMemoryStorage) Get(ctx context.Context, key string) (string, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	elem, ok := mem.store[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	mem.lru.MoveToFront(elem)
	return elem.Value.(MapEntry).value, nil
}
