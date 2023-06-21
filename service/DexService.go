package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/constant"
	c_factory "interest-arbitrage/contract/factory"
	"interest-arbitrage/global"
	"interest-arbitrage/model"
	"interest-arbitrage/model/query"
	"interest-arbitrage/server"
)

// /path/to/abigen --abi=/path/to/YourContractABI.json --pkg=main --out=/path/to/YourContract.go

func GetPair(c *gin.Context) {
	pairQuery := query.TokenPairQuery{}
	err := c.ShouldBindJSON(&pairQuery)
	if err != nil {
		model.Fail(c, model.ParamBindError)
		return
	}

	factoryAddress := common.HexToAddress(global.ContractInfoMap[constant.Factory].Address)

	contract, err := c_factory.NewContract(factoryAddress, server.UsableNodeServer.Client)
	key1 := common.HexToAddress(pairQuery.TokenA)
	key2 := common.HexToAddress(pairQuery.TokenB)
	result, err := contract.PairMap(&bind.CallOpts{}, key1, key2)
	if err != nil {
		model.Fail(c, model.CommonError, err.Error())
		return
	}

	model.Success(c, result.String())
}
