package service

import (
	"context"

	"github.com/dmsRosa6/go-shorty/internal/config"
	"github.com/dmsRosa6/go-shorty/internal/storage"
	"github.com/dmsRosa6/go-shorty/pkg/utils"
)

type URLShortenerService struct {
	store storage.Storage
}

func NewURLShortenerService() *URLShortenerService {
	return &URLShortenerService{
		store: storage.NewCompositeStorage(config.InMemoryCacheSize, config.RedisAddr, config.RedisDefaultExpiry),
	}
}

func (s *URLShortenerService) Shorten(ctx context.Context, url string) (string, error) {
	shorten := utils.TruncateSHA256Hex(url)

	err := s.store.Set(ctx, shorten, url)

	if err != nil {
		return "", err
	}

	return shorten, nil
}

func (s *URLShortenerService) Resolve(ctx context.Context, short string) (string, error) {
	return s.store.Get(ctx, short)
}
