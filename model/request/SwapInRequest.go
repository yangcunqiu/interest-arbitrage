package request

type SwapIn struct {
	TokenInAddress  string `json:"tokenInAddress,omitempty" binding:"required,IsAddress"`
	TokenOutAddress string `json:"tokenOutAddress,omitempty" binding:"required,IsAddress"`
	AmountIn        uint   `json:"amountIn,omitempty" binding:"required,gt=0"`
	AmountOut       uint   `json:"amountOut,omitempty" binding:"required,gt=0"`
	Spot            uint   `json:"spot,omitempty" binding:"required,gt=0,lt=100"`
}
