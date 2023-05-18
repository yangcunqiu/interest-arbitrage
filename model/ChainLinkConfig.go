package model

import "gorm.io/gorm"

type ChainLinkConfig struct {
	gorm.Model
	pairName      string `gorm:"type:varchar(20);not null;comment:价格对"`
	contractAddr  string `gorm:"type:char(42);not null;comment:合约地址"`
	abiJson       string `gorm:"type:json;not null;comment:合约abi"`
	priceFuncName string `gorm:"type:varchar(20);not null;comment:获取价格的函数名称"`
}

func (c ChainLinkConfig) tableName() string {
	return "chain_link_config"
}
