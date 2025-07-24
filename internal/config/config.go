package config

import "time"

const (
	InMemoryCacheSize        = 100
	RedisAddr                = "localhost:6379"
	RedisDefaultExpiry       = 10 * time.Minute
	TruncatedHashDefaultSize = 8
)
