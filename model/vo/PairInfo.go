package vo

import "math/big"

type PairInfo struct {
	Token0      string   `json:"token0,omitempty"`
	Token1      string   `json:"token1,omitempty"`
	Reserves0   *big.Int `json:"reserves0,omitempty"`
	Reserves1   *big.Int `json:"reserves1,omitempty"`
	TotalSupply *big.Int `json:"totalSupply,omitempty"`
}
