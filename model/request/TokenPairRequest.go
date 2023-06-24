package request

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
	"interest-arbitrage/utils"
)

type TokenPairRequest struct {
	TokenA string `json:"tokenA" binding:"required,IsAddress"`
	TokenB string `json:"tokenB" binding:"required,IsAddress"`
}

func IsAddress(fl validator.FieldLevel) bool {
	if address, ok := fl.Field().Interface().(string); ok {
		return common.IsHexAddress(address) && !utils.IsZeroAddress(address)
	}
	return false
}
