package entity

import (
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type UniswapV2PriceConfig struct {
	gorm.Model
	PairAddress   string `gorm:"type:char(42);default null;comment:pair地址"`
	Decimals      int    `gorm:"type:int;not null;comment:价格精度"`
	Token0Name    string `gorm:"type:varchar(50);not null;comment:token0名称"`
	Token0Address string `gorm:"type:char(42);not null;comment:token0地址"`
	Token1Name    string `gorm:"type:varchar(50);not null;comment:token1名称"`
	Token1Address string `gorm:"type:char(42);not null;comment:token1地址"`
	ShowToken     uint8  `gorm:"type:tinyint(1);comment:(0: token0, 1: token1)"`
}

func (c UniswapV2PriceConfig) TableName() string {
	return "uniswap_v2_price_config"
}

func ListUniswapPriceConfig() []UniswapV2PriceConfig {
	list := make([]UniswapV2PriceConfig, 0)
	global.DB.Find(&list)
	return list
}

func UpdateUniswapPriceConfig(config *UniswapV2PriceConfig) {
	global.DB.Model(config).Updates(config)
}
