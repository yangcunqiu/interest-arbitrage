package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"interest-arbitrage/model/request"
	"interest-arbitrage/utils"
	"math/big"
)

// AddLiquidity 添加流动性
func AddLiquidity(c *gin.Context) {
	addLiquidity := request.AddLiquidityRequest{}
	err := c.ShouldBindJSON(&addLiquidity)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 根据滑点计算能接受的最少数量
	amountAMin := addLiquidity.AmountADesired * (1 - addLiquidity.Spot/100.0)
	amountBMin := addLiquidity.AmountBDesired * (1 - addLiquidity.Spot/100.0)

	// 获取dex合约
	dex, err := utils.GetDexContract()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	// 构建opts
	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	// 发送交易
	tx, err := dex.AddLiquidity(
		opts,
		common.HexToAddress(addLiquidity.TokenA),
		common.HexToAddress(addLiquidity.TokenB),
		big.NewInt(int64(addLiquidity.AmountADesired)),
		big.NewInt(int64(addLiquidity.AmountBDesired)),
		big.NewInt(int64(amountAMin)),
		big.NewInt(int64(amountBMin)),
	)

	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

// SwapExactTokenForTokens 给定一个token的输入, 兑换出另一个token
func SwapExactTokenForTokens(c *gin.Context) {
	swapIn := request.SwapIn{}
	err := c.ShouldBindJSON(&swapIn)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 计算兑换路径
	swapPaths := getSwapPath(swapIn.TokenInAddress, swapIn.TokenOutAddress)
	// 根据滑点计算出最少接收数量
	amountOutMin := swapIn.AmountOut * (1 - swapIn.Spot/100.0)

	// 获取dex合约
	dex, err := utils.GetDexContract()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	// 构建opts
	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	tx, err := dex.SwapExactTokenForTokens(
		opts,
		big.NewInt(int64(swapIn.AmountIn)),
		big.NewInt(int64(amountOutMin)),
		swapPaths,
		common.HexToAddress(global.Env.Address),
	)

	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

// TODO 计算可兑换path
func getSwapPath(startAddress string, endAddress string) []common.Address {
	swapPaths := make([]common.Address, 2)
	swapPaths[0] = common.HexToAddress(startAddress)
	swapPaths[1] = common.HexToAddress(endAddress)
	return swapPaths
}

// RemoveLiquidity 移除流动性
func RemoveLiquidity(c *gin.Context) {
	removeReq := request.RemoveLiquidityRequest{}
	err := c.ShouldBindJSON(&removeReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	dex, err := utils.GetDexContract()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	tx, err := dex.RemoveLiquidity(
		opts,
		common.HexToAddress(removeReq.TokenA),
		common.HexToAddress(removeReq.TokenB),
		big.NewInt(int64(removeReq.Liquidity)),
		big.NewInt(int64(removeReq.AmountAMin)),
		big.NewInt(int64(removeReq.AMountBMin)),
		common.HexToAddress(removeReq.TokenTo),
	)

	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, tx.Hash())
}
