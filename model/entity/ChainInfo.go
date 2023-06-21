package entity

import (
	"errors"
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type ChainInfo struct {
	gorm.Model
	ChainId   int    `gorm:"type:int;not null;unique;comment:链id"`
	ChainName string `gorm:"type:varchar(50);not null;comment:链名称"`
	Currency  string `gorm:"type:varchar(50);not null;comment:原生币"`
	Type      int    `gorm:"type:int;not null;comment:类型(0 主网, 1 测试网)"`
}

func (c ChainInfo) TableName() string {
	return "chain_info"
}

func GetChainInfoById(id int) *ChainInfo {
	r := &ChainInfo{}
	global.DB.First(r, id)
	return r
}

func ListChainInfo() []ChainInfo {
	list := make([]ChainInfo, 0)
	global.DB.Find(&list)
	return list
}

func GetChainInfoByChainId(chainId int) *ChainInfo {
	r := &ChainInfo{}
	err := global.DB.Where("chain_id = ?", chainId).First(r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return r
}
