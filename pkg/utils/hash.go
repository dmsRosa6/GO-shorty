package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/dmsRosa6/go-shorty/internal/config"
)

//important reminder that the string on the result will be double the size of the bytes slices

func TruncateSHA256(s string) []byte {
	first := sha256.New()
	first.Write([]byte(s))
	hash := first.Sum(nil)

	return hash[:config.TruncatedHashDefaultSize]
}

func TruncateSHA256Hex(s string) string {
	hash := TruncateSHA256(s)
	return hex.EncodeToString(hash)
}

func TruncateSHA256WithSize(s string, size int) []byte {

	if size > sha256.Size {
		size = sha256.Size
	}

	first := sha256.New()
	first.Write([]byte(s))
	hash := first.Sum(nil)

	return hash[:size]
}

func TruncateSHA256HexWithSize(s string, size int) string {

	if size > sha256.Size {
		size = sha256.Size
	}

	hash := TruncateSHA256WithSize(s, size)
	return hex.EncodeToString(hash)
}
