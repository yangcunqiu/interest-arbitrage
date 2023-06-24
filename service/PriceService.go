package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/model"
	"interest-arbitrage/model/request"
	"interest-arbitrage/utils"
	"math/big"
)

// Quote 报价
func Quote(c *gin.Context) {
	quoteReq := request.QuoteRequest{}
	err := c.ShouldBindJSON(&quoteReq)
	if err != nil {
		model.Fail(c, model.ParamBindError)
		return
	}

	dex, err := utils.GetDexContract()
	if err != nil {
		model.Fail(c, model.GetContractError)
		return
	}

	quote, err := dex.Quote(
		&bind.CallOpts{},
		big.NewInt(int64(quoteReq.AmountIn)),
		big.NewInt(int64(quoteReq.ReserveIn)),
		big.NewInt(int64(quoteReq.ReserveOut)),
	)
	if err != nil {
		model.Fail(c, model.CallContractError)
		return
	}

	model.Success(c, quote)
}

// GetAmountOut 获取指定token能兑换出另一种token的数量
func GetAmountOut(c *gin.Context) {
	outReq := request.AmountOutRequest{}
	err := c.ShouldBindJSON(&outReq)
	if err != nil {
		model.Fail(c, model.ParamBindError)
		return
	}

	out, err := calAmountOut(outReq.AmountIn, outReq.ReserveIn, outReq.ReserveOut)
	if err != nil {
		model.Fail(c, model.CommonError, err.Error())
		return
	}

	model.Success(c, out)
}

func calAmountOut(amountIn uint, reserveIn uint, reserveOut uint) (*big.Int, error) {
	dex, err := utils.GetDexContract()
	if err != nil {
		return nil, err
	}

	out, err := dex.GetAmountOut(
		&bind.CallOpts{},
		big.NewInt(int64(amountIn)),
		big.NewInt(int64(reserveIn)),
		big.NewInt(int64(reserveOut)),
	)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// GetAmountIn 获取兑换出指定token数量需要的另一种token的数量
func GetAmountIn(c *gin.Context) {
	inReq := request.AmountInRequest{}
	err := c.ShouldBindJSON(&inReq)
	if err != nil {
		model.Fail(c, model.ParamBindError)
		return
	}

	dex, err := utils.GetDexContract()
	if err != nil {
		model.Fail(c, model.GetContractError)
		return
	}

	in, err := dex.GetAmountIn(
		&bind.CallOpts{},
		big.NewInt(int64(inReq.AmountOut)),
		big.NewInt(int64(inReq.ReserveIn)),
		big.NewInt(int64(inReq.ReserveOut)),
	)
	if err != nil {
		model.Fail(c, model.CallContractError)
		return
	}

	model.Success(c, in)
}
