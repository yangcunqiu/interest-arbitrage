package request

type AmountInRequest struct {
	AmountOut  uint `json:"amountOut,omitempty" binding:"gt=0"`
	ReserveIn  uint `json:"reserveIn,omitempty" binding:"gt=0"`
	ReserveOut uint `json:"reserveOut,omitempty" binding:"gt=0"`
}
