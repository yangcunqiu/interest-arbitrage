package entity

import (
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type BridgeLockInfo struct {
	gorm.Model
	ChainId      int    `gorm:"type:int;not null;comment:链id"`
	UserAddress  string `gorm:"type:char(42);not null;comment:用户地址"`
	TokenAddress string `gorm:"type:char(42);comment:token地址"`
	TokenType    uint   `gorm:"type:int;not null;comment:token类型, 0: 原生代币, 1: erc20, 2: erc721"`
	TokenSymbol  string `gorm:"type:varchar(20);not null;comment:token简称"`
	Amount       string `gorm:"type:varchar(200);not null;comment:金额"`
	Decimal      uint   `gorm:"type:int;not null;comment:精度"`
	Unit         string `gorm:"type:varchar(20);not null;comment:单位"`
}

func (bli BridgeLockInfo) TableName() string {
	return "bridge_lock_info"
}

func GetAllBridgeLockInfo() []*BridgeLockInfo {
	results := make([]*BridgeLockInfo, 0)
	global.DB.Model(&BridgeLockInfo{}).Find(&results)
	return results
}

func GetBridgeLockInfoByChainId(chainId int) []*BridgeLockInfo {
	results := make([]*BridgeLockInfo, 0)
	global.DB.Model(&BridgeLockInfo{}).Where("chain_id = ?", chainId).Find(&results)
	return results
}

func GetBridgeLockInfoByUserAndToken(userAddress string, tokenAddress string) (*BridgeLockInfo, error) {
	info := &BridgeLockInfo{}
	err := global.DB.Model(&BridgeLockInfo{}).Where("user_address = ? and token_address = ?", userAddress, tokenAddress).First(info).Error
	return info, err
}

func SaveBridgeLockInfo(bli *BridgeLockInfo) error {
	return global.DB.Create(bli).Error
}

func UpdateBridgeLockInfoById(bli *BridgeLockInfo) error {
	return global.DB.Updates(bli).Error
}
