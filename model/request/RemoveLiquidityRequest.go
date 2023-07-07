package request

type RemoveLiquidityRequest struct {
	TokenA     string `json:"tokenA,omitempty" binding:"required,IsAddress"`
	TokenB     string `json:"tokenB,omitempty" binding:"required,IsAddress"`
	Liquidity  uint   `json:"liquidity,omitempty" binding:"required,gt=0"`
	AmountAMin uint   `json:"amountAMin,omitempty" binding:"required,gt=0"`
	AMountBMin uint   `json:"AMountBMin,omitempty" binding:"required,gt=0"`
	TokenTo    string `json:"tokenTo,omitempty" binding:"required,IsAddress"`
}
