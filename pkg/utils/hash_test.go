package utils_test

import (
	"testing"

	"github.com/dmsRosa6/go-shorty/pkg/utils"
	"github.com/stretchr/testify/require"
)

func Test_Hashing(t *testing.T) {
	text := "Hello"
	hash := "185f8db322"
	size := 5

	res := utils.TruncateSHA256HexWithSize(text, size)
	require.Equal(t, res, hash)
}

func Test_Hashing_DifferentSize(t *testing.T) {
	text := "Hello"
	size := 3
	res := utils.TruncateSHA256HexWithSize(text, size)
	require.Len(t, res, 6)
}

func Test_Hashing_ZeroSize(t *testing.T) {
	res := utils.TruncateSHA256HexWithSize("Hello", 0)
	require.Equal(t, "", res)
}

// This will not panic because i cap the max size to the sha264 max size
func Test_Hashing_Oversized(t *testing.T) {
	text := "Hello"
	size := 40
	res := utils.TruncateSHA256HexWithSize(text, size)

	require.Equal(t, len(res), 64)
}
