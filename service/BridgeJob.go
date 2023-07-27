package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"interest-arbitrage/constant"
	"interest-arbitrage/contract/eb"
	"interest-arbitrage/contract/mb"
	"interest-arbitrage/global"
	"interest-arbitrage/model/entity"
	"interest-arbitrage/server"
	"interest-arbitrage/utils"
	"log"
)

func SyncETHBridgeLockMtsByAll() {
	ethBridge, err := utils.GetETHBridge()
	if err != nil {
		log.Fatalf("failed to get ETHBridge contract, err: %v", err)
		return
	}

	// get all user lock info
	lockInfos := entity.GetBridgeLockInfoByChainId(int(server.ETHNodeServer.ChainId.Int64()))
	for _, lockInfo := range lockInfos {
		syncETHBridgeMtsLock(ethBridge, lockInfo)
	}
	log.Println("success sync ETHBridge user lock mts")
}

func SyncETHBridgeLockMtsByAccount(account string) {
	ethBridge, err := utils.GetETHBridge()
	if err != nil {
		log.Fatalf("failed to get ETHBridge contract, err: %v", err)
		return
	}
	lockInfo, _ := entity.GetBridgeLockInfoByUserAndToken(account, global.ContractInfoMap[constant.ETHBridge].Address)
	if lockInfo != nil {
		syncETHBridgeMtsLock(ethBridge, lockInfo)
	}
	log.Printf("success sync ETHBridge user lock mts by account: %v", account)
}

func syncETHBridgeMtsLock(ethBridge *eb.ETHBridge, lockInfo *entity.BridgeLockInfo) {
	amount, err := ethBridge.MtsAmountMap(&bind.CallOpts{}, common.HexToAddress(lockInfo.UserAddress))
	if err != nil {
		log.Printf("failed to get mtsAmountMap in ETHBridge, err: %v", err)
	}
	lockInfo.Amount = amount.String()
	err = entity.UpdateBridgeLockInfoById(lockInfo)
	if err != nil {
		log.Printf("failed to updateUserLockInfo, err: %v", err)
	}
}

func SyncMTSBridgeLockMtsByAll() {
	mtsBridge, err := utils.GetMTSBridge()
	if err != nil {
		log.Fatalf("failed to get MTSBridge contract, err: %v", err)
		return
	}

	// get all user lock info
	lockInfos := entity.GetBridgeLockInfoByChainId(int(server.MTSNodeServer.ChainId.Int64()))
	for _, lockInfo := range lockInfos {
		syncMTSBridgeLockMts(mtsBridge, lockInfo)
	}
	log.Println("success sync MTSBridge user lock mts")
}

func SyncMTSBridgeLockMtsByAccount(account string) {
	mtsBridge, err := utils.GetMTSBridge()
	if err != nil {
		log.Fatalf("failed to get MTSBridge contract, err: %v", err)
		return
	}
	lockInfo, _ := entity.GetBridgeLockInfoByUserAndToken(account, global.ContractInfoMap[constant.MTSBridge].Address)
	if lockInfo != nil {
		syncMTSBridgeLockMts(mtsBridge, lockInfo)
	}
	log.Printf("success sync MTSBridge user lock mts by account: %v", account)
}

func syncMTSBridgeLockMts(mtsBridge *mb.MTSBridge, lockInfo *entity.BridgeLockInfo) {
	amount, err := mtsBridge.MtsAmountMap(&bind.CallOpts{}, common.HexToAddress(lockInfo.UserAddress))
	if err != nil {
		log.Printf("failed to get mtsAmountMap in MTSBridge, err: %v", err)
	}
	lockInfo.Amount = amount.String()
	err = entity.UpdateBridgeLockInfoById(lockInfo)
	if err != nil {
		log.Printf("failed to updateUserLockInfo, err: %v", err)
	}
}
