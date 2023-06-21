package global

import (
	"gorm.io/gorm"
	"interest-arbitrage/config"
	"interest-arbitrage/model"
)

var (
	DB              *gorm.DB
	Config          config.Config
	Env             config.Env
	ContractInfoMap map[string]*model.ContractInfo
)
