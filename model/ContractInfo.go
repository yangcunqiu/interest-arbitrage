package model

import "github.com/ethereum/go-ethereum/accounts/abi"

type ContractInfo struct {
	Name    string
	Address string
	ABIPath string
	ABI     *abi.ABI
}
