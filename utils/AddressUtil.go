package utils

import (
	"bytes"
	"interest-arbitrage/constant"
)

func IsZeroAddress(address string) bool {
	return bytes.Equal([]byte(address), []byte(constant.ZeroAddress))
}
