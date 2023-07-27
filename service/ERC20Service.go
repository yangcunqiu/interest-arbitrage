package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"interest-arbitrage/model/request"
	"interest-arbitrage/server"
	"interest-arbitrage/utils"
	"math/big"
)

// BalanceOf 查看余额
func BalanceOf(c *gin.Context) {
	balanceReq := request.BalanceRequest{}
	err := c.ShouldBindJSON(&balanceReq)
	if err != nil {
		model.Fail(c, model.ParamBindError)
		return
	}

	erc20, err := utils.GetERC20(balanceReq.ERC20TokenAddress)
	if err != nil {
		model.Fail(c, model.GetContractError)
		return
	}

	balanceOf, err := erc20.BalanceOf(&bind.CallOpts{}, common.HexToAddress(balanceReq.AccountAddress))
	if err != nil {
		model.Fail(c, model.CallContractError)
		return
	}

	model.Success(c, balanceOf)
}

// Approve 授权MyDex合约erc20token
func Approve(c *gin.Context) {
	approve := request.TokenApprove{}
	err := c.ShouldBindJSON(&approve)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 授权
	erc20, err := utils.GetERC20(approve.ERC20TokenAddress)
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(server.UsableNodeServer.ChainId, global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError)
		return
	}
	tx, err := erc20.Approve(
		opts,
		common.HexToAddress(approve.SpenderAddress),
		big.NewInt(int64(approve.Amount)),
	)

	if err != nil {
		model.Fail(c, model.CallContractError)
		return
	}

	// 只是交易hash, 此时交易是未被确认的
	model.Success(c, tx.Hash())
}
