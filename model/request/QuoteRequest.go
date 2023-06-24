package request

type QuoteRequest struct {
	AmountIn   uint `json:"amountIn,omitempty" binding:"gt=0"`
	ReserveIn  uint `json:"reserveIn,omitempty" binding:"gt=0"`
	ReserveOut uint `json:"reserveOut,omitempty" binding:"gt=0"`
}
