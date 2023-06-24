package test

import (
	"fmt"
	"log"
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	bigA := big.NewInt(20)
	bigB := big.NewInt(15)
	bigC := big.NewInt(3)
	result := new(big.Int)
	result.Sub(bigA, bigB).Div(result, bigC)
	log.Println(result)

	fA := new(big.Float).SetInt(bigA)
	fB := new(big.Float).SetInt(bigB)
	sub := new(big.Float).Sub(fB, fA)
	quo := new(big.Float).Quo(sub, fA)
	resultA := quo.Mul(quo, big.NewFloat(100))
	sprintf := fmt.Sprintf("%.2f%%", resultA)
	log.Println(sprintf)
}

func TestFloat(t *testing.T) {
	fmt.Println(100 * (1 - 1/100.0))
}
