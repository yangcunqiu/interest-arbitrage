package model

import (
	"gorm.io/gorm"
	"math/big"
)

type UniswapV2TokenCumulativeLast struct {
	gorm.Model
	PairAddress          string
	Price0CumulativeLast big.Int
	Price1CumulativeLast big.Int
	BlockTimestampLast   big.Int
}
