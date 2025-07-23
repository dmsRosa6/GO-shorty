package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// TODO: Move this to a config file or constan
const size = 10

func TruncateSHA256(s string) []byte {
	first := sha256.New()
	first.Write([]byte(s))
	hash := first.Sum(nil)

	return hash[:size]
}

func TruncateSHA256Hex(s string) string {
	hash := TruncateSHA256(s)
	return hex.EncodeToString(hash)
}
