package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"interest-arbitrage/model/request"
	"interest-arbitrage/utils"
	"math/big"
	"strconv"
)

// AddAllot 添加一种lpToken
func AddAllot(c *gin.Context) {
	addAllotReq := request.AddLpTokenAllotRequest{}
	err := c.ShouldBindJSON(&addAllotReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	tx, err := allot.Add(
		opts,
		common.HexToAddress(addAllotReq.LpToken),
		big.NewInt(int64(addAllotReq.AllocPoint)),
		addAllotReq.IsUpdate,
	)

	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, tx.Hash())
}

// GetPending 查看用户待领取奖励
func GetPending(c *gin.Context) {
	account := c.Query("account")
	pid, _ := strconv.Atoi(c.Query("pid"))

	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	pendingAmount, err := allot.PendingMdt(&bind.CallOpts{}, common.HexToAddress(account), big.NewInt(int64(pid)))
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, pendingAmount)
}

// Deposit 存lp
func Deposit(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	amount, _ := strconv.Atoi(c.Query("amount"))

	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	tx, err := allot.Deposit(opts, big.NewInt(int64(pid)), big.NewInt(int64(amount)))
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, tx.Hash())
}

// Withdraw 取出lp
func Withdraw(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	amount, _ := strconv.Atoi(c.Query("amount"))

	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	tx, err := allot.Withdraw(opts, big.NewInt(int64(pid)), big.NewInt(int64(amount)))
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, tx.Hash())
}
