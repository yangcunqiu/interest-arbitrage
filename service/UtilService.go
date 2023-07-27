package service

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"interest-arbitrage/model/request"
	"interest-arbitrage/model/vo"
	"interest-arbitrage/server"
	"interest-arbitrage/utils"
	"math/big"
	"strconv"
	"time"
)

// /path/to/abigen --abi=/path/to/YourContractABI.json --pkg=main --out=/path/to/YourContract.go

/*
Receipt
root: 表示状态根哈希，表示交易执行后的状态树根哈希。如果状态树没有发生变化，则该值为 0x。
status: 表示交易执行的状态。如果值为 0x1，表示交易成功执行；其他非零值或错误表示交易执行失败。
cumulativeGasUsed: 表示该交易及之前的交易在区块中累计消耗的燃气总量。
logsBloom: 日志布隆过滤器，用于快速检索与该交易关联的日志事件。
logs: 一个包含与该交易关联的日志事件的数组。日志事件记录了合约中定义的事件的具体信息。
transactionHash: 交易的哈希值，唯一标识该交易。
contractAddress: 如果该交易是合约创建交易，该字段表示新创建的合约地址；否则为 0x。
gasUsed: 表示该交易消耗的燃气量。
effectiveGasPrice: 表示该交易的有效燃气价格。
blockHash: 区块的哈希值，表示该交易所在的区块。
blockNumber: 区块的编号，表示该交易所在的区块高度。
transactionIndex: 交易在区块中的索引位置。
*/

// GetPair 获取token池子地址
func GetPair(c *gin.Context) {
	// 参数绑定
	pairQuery := request.TokenPairRequest{}
	err := c.ShouldBindJSON(&pairQuery)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	pair, err := getPair(pairQuery.TokenA, pairQuery.TokenB)
	if err != nil {
		model.Fail(c, model.CommonError, err.Error())
		return
	}

	model.Success(c, pair.String())
}

func getPair(tokenA string, tokenB string) (*common.Address, error) {
	// 获取factory合约
	factory, err := utils.GetFactoryContract()
	if err != nil {
		return nil, err
	}

	// 查询
	pair, err := factory.PairMap(
		&bind.CallOpts{},
		common.HexToAddress(tokenA),
		common.HexToAddress(tokenB),
	)
	if err != nil {
		return nil, err
	}
	return &pair, nil
}

// GetPairInfo 获取pair信息
func GetPairInfo(c *gin.Context) {
	pairAddress := c.Query("pairAddress")
	if pairAddress == "" || !common.IsHexAddress(pairAddress) {
		model.Fail(c, model.AddressNotBlankOrNotAddress)
		return
	}

	pair, err := utils.GetPair(pairAddress)
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	token0, err := pair.Token0(&bind.CallOpts{})
	token1, err := pair.Token1(&bind.CallOpts{})
	reserves0, err := pair.Reserves0(&bind.CallOpts{})
	reserves1, err := pair.Reserves1(&bind.CallOpts{})
	totalSupply, err := pair.TotalSupply(&bind.CallOpts{})
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	model.Success(c, vo.PairInfo{
		Token0:      token0.String(),
		Token1:      token1.String(),
		Reserves0:   reserves0,
		Reserves1:   reserves1,
		TotalSupply: totalSupply,
	})
}

func getReserves(pairAddress string) (reserves0 *big.Int, reserves1 *big.Int, err error) {
	reserves0 = nil
	reserves1 = nil
	pair, err := utils.GetPair(pairAddress)
	if err != nil {
		return
	}

	reserves0, err = pair.Reserves0(&bind.CallOpts{})
	reserves1, err = pair.Reserves1(&bind.CallOpts{})
	if err != nil {
		return
	}
	return
}

// CreatePair 创建池子
func CreatePair(c *gin.Context) {
	addPair := request.TokenPairRequest{}
	err := c.ShouldBindJSON(&addPair)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 获取factory合约
	factory, err := utils.GetFactoryContract()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	opts, err := utils.BuildTransactOpts(server.UsableNodeServer.ChainId, global.Env.PrivateKey)
	if err != nil {
		model.Fail(c, model.BuildTransactOptsError, err.Error())
		return
	}

	// 如果err=nil只是代表这个交易被发送到区块链中, 但是还未被确认, 要等待矿工打包
	tx, err := factory.CreatePair(
		opts,
		common.HexToAddress(addPair.TokenA),
		common.HexToAddress(addPair.TokenB),
	)
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}

	// 等待交易确认
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()

	receipt, err := bind.WaitMined(ctx, server.UsableNodeServer.Client, tx)
	if err != nil {
		if err == context.DeadlineExceeded {
			model.Success(c, model.WaitTransactionPackTimeOut.Message, tx.Hash().String())
			return
		} else {
			model.Success(c, model.WaitTransactionPackError.Message, tx.Hash().String())
			return
		}
	}

	model.Success(c, receipt)
}

// GetPoolInfoList 查看MyDexTokenAllot合约所有lp池
func GetPoolInfoList(c *gin.Context) {
	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}

	var pools []vo.PoolVO
	for i := 0; ; i++ {
		poolInfo, err := allot.PoolInfos(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			break
		}

		if utils.IsZeroAddress(poolInfo.LpToken.String()) {
			break
		}

		pool := vo.PoolVO{
			Pid:     uint(i),
			LpToken: poolInfo.LpToken.String(),
		}
		pools = append(pools, pool)
	}

	model.Success(c, pools)
}

// GetPoolInfo 获取lp池信息
func GetPoolInfo(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	allot, err := utils.GetDexTokenAllot()
	if err != nil {
		model.Fail(c, model.GetContractError, err.Error())
		return
	}
	poolInfo, err := allot.PoolInfos(&bind.CallOpts{}, big.NewInt(int64(pid)))
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, poolInfo)
}
