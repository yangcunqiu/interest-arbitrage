package service

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"interest-arbitrage/constant"
	"interest-arbitrage/contract/eb"
	"interest-arbitrage/contract/mb"
	"interest-arbitrage/global"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"interest-arbitrage/model/entity"
	"interest-arbitrage/model/request"
	"interest-arbitrage/server"
	"interest-arbitrage/utils"
	"log"
	"math/big"
)

func ListenerMTSBridgeLockMTS() {
	mtsBridge, err := utils.GetMTSBridge()
	if err != nil {
		log.Printf("failed to get MTSBridge contract, err: %v", err)
	}

	eventCh := make(chan *mb.MTSBridgeLockMts)
	sub, err := mtsBridge.WatchLockMts(&bind.WatchOpts{}, eventCh, nil)
	if err != nil {
		log.Printf("failed to watch event MTSBridge lockMts, err: %v", err)
	}

	go func() {
		for {
			select {
			case err = <-sub.Err():
				// log.Printf("failed to sub event MTSBridge lockMts, err: %v", err)
			case vLog := <-eventCh:
				log.Printf("success to sub event MTSBridge lockMts, vLog: %v", vLog)
				SyncMTSBridgeLockMtsByAccount(vLog.Account.String())
			}
		}
	}()
}

func ListenerETHBridgeLockMts() {
	ethBridge, err := utils.GetETHBridge()
	if err != nil {
		log.Printf("failed to get ETHBridge contract, err: %v", err)
	}

	eventCh := make(chan *eb.ETHBridgeLockMts)
	sub, err := ethBridge.WatchLockMts(&bind.WatchOpts{}, eventCh, nil)
	if err != nil {
		log.Printf("failed to watch event ETHBridge lockMts, err: %v", err)
	}

	go func() {
		for {
			select {
			case err = <-sub.Err():
				log.Printf("failed to sub event ETHBridge lockMts, err: %v", err)
			case vLog := <-eventCh:
				log.Printf("success to sub event ETHBridge lockMts, vLog: %v", vLog)
				SyncETHBridgeLockMtsByAccount(vLog.Account.String())
			}
		}
	}()
}

func updateUserLockInfo(chainId int, userAddress string, tokenAddress string, tokenType uint, tokenSymbol string, amount string, decimal uint, unit string) error {
	bridgeLockInfo, err := entity.GetBridgeLockInfoByUserAndToken(userAddress, tokenAddress)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 新增
		lockInfo := &entity.BridgeLockInfo{
			ChainId:      chainId,
			UserAddress:  userAddress,
			TokenAddress: tokenAddress,
			TokenType:    tokenType,
			TokenSymbol:  tokenSymbol,
			Amount:       amount,
			Decimal:      decimal,
			Unit:         unit,
		}
		return entity.SaveBridgeLockInfo(lockInfo)
	}

	if err != nil {
		log.Printf("failed to get bridgeLockInfo, err: %v", err)
		return err
	}

	// 修改
	bigA := new(big.Int)
	bigB := new(big.Int)
	_, _ = bigA.SetString(bridgeLockInfo.Amount, 10)
	_, _ = bigB.SetString(amount, 10)
	bigC := new(big.Int)
	bigC.Add(bigA, bigB)

	bridgeLockInfo.Amount = bigC.String()
	return entity.UpdateBridgeLockInfoById(bridgeLockInfo)
}

func LockMtsInMTS(c *gin.Context) {
	lockReq := request.LockMtsRequest{}
	err := c.ShouldBindJSON(&lockReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 记录lockInfo, 但是不记录amount, amount只能由event和job更新, 数据来源必须是合约
	_, err = entity.GetBridgeLockInfoByUserAndToken(lockReq.UserAddress, global.ContractInfoMap[constant.MTSBridge].Address)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 新增
		lockInfo := &entity.BridgeLockInfo{
			ChainId:      int(server.MTSNodeServer.ChainId.Int64()),
			UserAddress:  lockReq.UserAddress,
			TokenAddress: global.ContractInfoMap[constant.MTSBridge].Address,
			TokenType:    0,
			TokenSymbol:  "mts",
			Decimal:      18,
			Unit:         "mts",
		}
		err := entity.SaveBridgeLockInfo(lockInfo)
		if err != nil {
			model.Fail(c, model.LockMtsError, err.Error())
			return
		}
	}

	mtsBridge, _ := utils.GetMTSBridge()
	opts, _ := utils.BuildTransactOpts(server.MTSNodeServer.ChainId, global.Env.MTSPrivateKey1)

	ethValue := big.NewFloat(10) // 10mts
	weiValue := new(big.Int)
	// 将 ETH 转换为 Wei
	ethValue.Mul(ethValue, new(big.Float).SetInt(big.NewInt(params.Ether)))
	weiValue, _ = ethValue.Int(weiValue)

	opts.Value = weiValue
	tx, err := mtsBridge.LockMts(opts)
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

func UnLockMtsInMTS(c *gin.Context) {
	unlockReq := request.UnlockMtsRequest{}
	err := c.ShouldBindJSON(&unlockReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 校验
	lockInfo, err := entity.GetBridgeLockInfoByUserAndToken(unlockReq.UserAddress, global.ContractInfoMap[constant.MTSBridge].Address)
	bigA := new(big.Int)
	bigB := new(big.Int)
	_, _ = bigA.SetString(lockInfo.Amount, 10)
	_, _ = bigB.SetString(unlockReq.Amount, 10)
	if errors.Is(err, gorm.ErrRecordNotFound) || lockInfo == nil || bigA.Cmp(bigB) < 0 {
		model.Fail(c, model.LockMtsNotEnough)
		return
	}

	mtsBridge, _ := utils.GetMTSBridge()
	opts, _ := utils.BuildTransactOpts(server.MTSNodeServer.ChainId, global.Env.MTSPrivateKey)
	tx, err := mtsBridge.UnlockMts(opts, common.HexToAddress(unlockReq.UserAddress), bigB)
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

func LockMtsInETH(c *gin.Context) {
	lockReq := request.LockMtsRequest{}
	err := c.ShouldBindJSON(&lockReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 记录lockInfo, 但是不记录amount, amount只能由event和job更新, 数据来源必须是合约
	_, err = entity.GetBridgeLockInfoByUserAndToken(lockReq.UserAddress, global.ContractInfoMap[constant.ETHBridge].Address)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 新增
		lockInfo := &entity.BridgeLockInfo{
			ChainId:      int(server.ETHNodeServer.ChainId.Int64()),
			UserAddress:  lockReq.UserAddress,
			TokenAddress: global.ContractInfoMap[constant.ETHBridge].Address,
			TokenType:    1,
			TokenSymbol:  "eth",
			Decimal:      18,
			Unit:         "mts",
		}
		err := entity.SaveBridgeLockInfo(lockInfo)
		if err != nil {
			model.Fail(c, model.LockMtsError, err.Error())
			return
		}
	}

	ethBridge, _ := utils.GetETHBridge()
	opts, _ := utils.BuildTransactOpts(server.ETHNodeServer.ChainId, global.Env.ETHPrivateKey1)

	bigA := new(big.Int)
	bigA.SetString(lockReq.Amount, 10)
	tx, err := ethBridge.LockMts(opts, bigA)
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

func UnLockMtsInETH(c *gin.Context) {
	unlockReq := request.UnlockMtsRequest{}
	err := c.ShouldBindJSON(&unlockReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	// 校验
	lockInfo, err := entity.GetBridgeLockInfoByUserAndToken(unlockReq.UserAddress, global.ContractInfoMap[constant.MTSBridge].Address)
	bigA := new(big.Int)
	bigB := new(big.Int)
	_, _ = bigA.SetString(lockInfo.Amount, 10)
	_, _ = bigB.SetString(unlockReq.Amount, 10)
	if errors.Is(err, gorm.ErrRecordNotFound) || lockInfo == nil || bigA.Cmp(bigB) < 0 {
		model.Fail(c, model.LockMtsNotEnough)
		return
	}

	ethBridge, _ := utils.GetETHBridge()
	opts, _ := utils.BuildTransactOpts(server.ETHNodeServer.ChainId, global.Env.ETHPrivateKey)
	tx, err := ethBridge.UnlockMts(opts, common.HexToAddress(unlockReq.UserAddress), bigB)
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}

func MintWMTS(c *gin.Context) {
	mintReq := request.MintRequest{}
	err := c.ShouldBindJSON(&mintReq)
	if err != nil {
		model.Fail(c, model.ParamBindError, handler.SimpleValidateErrorTran(err))
		return
	}

	erc20, err := utils.GetWMTS()
	if err != nil {
		model.Fail(c, model.GetContractError)
	}
	opts, _ := utils.BuildTransactOpts(server.ETHNodeServer.ChainId, global.Env.ETHPrivateKey)
	tx, err := erc20.Mint(opts, common.HexToAddress(mintReq.Account), big.NewInt(int64(mintReq.Amount)))
	if err != nil {
		model.Fail(c, model.CallContractError, err.Error())
		return
	}
	model.Success(c, tx.Hash())
}
