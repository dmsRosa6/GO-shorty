package config

import "time"

const (
	InMemoryCacheSize        = 100
	RedisAddr                = "localhost:6379"
	RedisDefaultExpiry       = 10 * time.Minute
	TruncatedHashDefaultSize = 8
	InDaClub                 = "https://www.youtube.com/watch?v=6v_TUnNnf2E"
	InDaClubProbability      = 1
)
