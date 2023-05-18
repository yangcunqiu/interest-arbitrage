package global

import (
	"gorm.io/gorm"
	"interest-arbitrage/config"
)

var (
	DB     *gorm.DB
	Config config.Config
	Env    config.Env
)
