package helpers

import (
	"crypto/rand"
	"math/big"
)

func RandomInt() int8 {
	bg := big.NewInt(100 - 1)

	res, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	return int8(res.Int64() + 1)
}
