package entity

import (
	"errors"
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type Price struct {
	gorm.Model
	TokenName     string `gorm:"type:varchar(50);not null;comment:token名称"`
	TokenAddress  string `gorm:"type:char(42);not null;unique;comment:token地址"`
	Value         string `gorm:"type:varchar(50);not null;comment:原始价格"`
	Decimals      int    `gorm:"type:int;not null;comment:精度"`
	ReadablePrice string `gorm:"type:varchar(50);not null;comment:易读价格"`
	Markup        string `gorm:"type:varchar(10);comment:涨幅"`
	Unit          string `gorm:"type:varchar(50);not null;comment:价格单位"`
	SourceType    int    `gorm:"type:int;not null;comment:来源类型(0 chainLink, 1 uniswap)"`
	ChainId       int    `gorm:"type:int;not null;comment:链id"`
}

func (p Price) TableName() string {
	return "price"
}

func SavePrice(price *Price) {
	global.DB.Model(&Price{}).Create(price)
}

func GetPriceByTokenAddress(tokenAddress string) *Price {
	p := &Price{}
	err := global.DB.Where("token_address = ?", tokenAddress).First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return p
}

func UpdatePrice(price *Price) {
	global.DB.Model(price).Updates(price)
}

func SaveOrUpdatePrice(price *Price) {
	byTokenAddress := GetPriceByTokenAddress(price.TokenName)
	if byTokenAddress == nil {
		SavePrice(price)
	} else {
		UpdatePrice(price)
	}
}
