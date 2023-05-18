package model

import (
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type ChainLinkConfig struct {
	gorm.Model
	PairName      string `gorm:"type:varchar(20);not null;comment:价格对"`
	ContractAddr  string `gorm:"type:char(42);not null;comment:合约地址"`
	AbiJson       string `gorm:"type:json;not null;comment:合约abi"`
	PriceFuncName string `gorm:"type:varchar(20);not null;comment:获取价格的函数名称"`
	Decimals      int    `gorm:"type:int;not null;comment:价格精度"`
	ChainId       int    `gorm:"type:int;comment:链id"`
	Type          string `gorm:"type:varchar(20);comment:类型"`
}

func (c ChainLinkConfig) tableName() string {
	return "chain_link_config"
}

func ListChainLinkConfig() []ChainLinkConfig {
	list := make([]ChainLinkConfig, 0)
	global.DB.Find(&list)
	return list
}
