package entity

import (
	"errors"
	"gorm.io/gorm"
	"interest-arbitrage/global"
)

type UniswapV2TokenCumulativeLast struct {
	gorm.Model
	PairAddress          string `gorm:"type:char(42);not null;comment:pair地址"`
	Price0Average        string `gorm:"type:varchar(100);comment:reserve0"`
	Price1Average        string `gorm:"type:varchar(100);comment:reserve0"`
	Reserve0             string `gorm:"type:varchar(100);not null;comment:reserve0"`
	Reserve1             string `gorm:"type:varchar(100);not null;comment:reserve1"`
	Price0CumulativeLast string `gorm:"type:varchar(100);not null;comment:Price0CumulativeLast"`
	Price1CumulativeLast string `gorm:"type:varchar(100);not null;comment:Price1CumulativeLast"`
	BlockTimestampLast   uint32 `gorm:"type:int;not null;comment:时间戳"`
}

func (u UniswapV2TokenCumulativeLast) TableName() string {
	return "uniswap_v2_token_cumulative_last"
}

func GetTokenCumulativeLastByPairAddress(pairAddress string) *UniswapV2TokenCumulativeLast {
	result := new(UniswapV2TokenCumulativeLast)
	tx := global.DB.Where("pair_address = ?", pairAddress).First(result)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return result
}

func SaveTokenCumulativeLast(tcl *UniswapV2TokenCumulativeLast) {
	global.DB.Model(&UniswapV2TokenCumulativeLast{}).Create(tcl)
}

func SaveOrUpdateTokenCumulativeLast(tcl *UniswapV2TokenCumulativeLast) {
	byPairAddress := GetTokenCumulativeLastByPairAddress(tcl.PairAddress)
	if byPairAddress == nil {
		SaveTokenCumulativeLast(tcl)
		return
	}
	global.DB.Model(tcl).Where("pair_address = ?", tcl.PairAddress).Updates(tcl)
}
