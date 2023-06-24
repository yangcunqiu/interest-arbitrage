package request

type AddLiquidityRequest struct {
	// tokenA地址
	TokenA string `json:"tokenA,omitempty" binding:"required,IsAddress"`
	// tokenB地址
	TokenB string `json:"tokenB,omitempty" binding:"required,IsAddress"`
	// tokenA期望添加数量
	AmountADesired uint `json:"amountADesired,omitempty" binding:"required,gt=0"`
	// tokenB期望添加数量
	AmountBDesired uint `json:"amountBDesired,omitempty" binding:"required,gt=0"`
	// 滑点%
	Spot uint `json:"spot,omitempty" binding:"required,gt=0,lt=100"`
}
