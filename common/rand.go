package common

import (
	"crypto/rand"
	"math/big"
)

func Rand(max uint) uint {
	maxBigInt := big.NewInt(int64(max - 1))
	i, _ := rand.Int(rand.Reader, maxBigInt)
	return uint(i.Int64() + 1)
}
