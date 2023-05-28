package test

import (
	"log"
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	bigA := big.NewInt(16)
	bigB := big.NewInt(6)
	bigC := big.NewInt(3)
	result := new(big.Int)
	result.Sub(bigA, bigB).Div(result, bigC)
	log.Println(result)
}
