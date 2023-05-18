package service

import (
	"interest-arbitrage/model"
	"math/big"
)

type PriceLister interface {
	getPrice(model.ChainLinkConfig) *big.Int
}
