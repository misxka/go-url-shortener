package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Encode(str string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(str))
	return algorithm.Sum(nil)
}

func base58Encode(bytes []byte) string {
	encoded, err := base58.BitcoinEncoding.Encode(bytes)

	if err != nil {
		panic(err)
	}

	return string(encoded)
}

func GenerateShortUrl(original string, userId string) string {
	hashed := sha256Encode(original + userId)
	numericId := new(big.Int).SetBytes(hashed).Uint64()
	return base58Encode(fmt.Appendf(nil, "%d", numericId))[:8]
}
