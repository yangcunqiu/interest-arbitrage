package entity

import (
	"errors"
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type TokenPrice struct {
	gorm.Model
	PairName  string `gorm:"type:varchar(20);not null;comment:价格对"`
	Decimals  int    `gorm:"type:int;not null;comment:价格精度"`
	Price     string `gorm:"type:varchar(500);not null;comment:价格"`
	ReadPrice string `gorm:"type:varchar(300);not null;comment:价格"`
	ChainId   int    `gorm:"type:int;comment:链id"`
	Type      string `gorm:"type:varchar(20);comment:类型"`
}

func (tp TokenPrice) TableName() string {
	return "token_price"
}

func SaveTokenPrice(tp *TokenPrice) {
	global.DB.Model(&TokenPrice{}).Create(tp)
}

func GetTokenPriceByPariName(pairName string) *TokenPrice {
	result := new(TokenPrice)
	tx := global.DB.Model(result).Where("pair_name = ?", pairName).First(result)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return result
}

func SaveOrUpdateTokenPrice(tp *TokenPrice) {
	byPariName := GetTokenPriceByPariName(tp.PairName)
	if byPariName != nil {
		global.DB.Model(&tp).Where("pair_name = ?", tp.PairName).Updates(&tp)
		return
	}
	SaveTokenPrice(tp)
}
