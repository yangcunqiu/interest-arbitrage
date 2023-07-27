package request

type MintRequest struct {
	Account string `json:"account,omitempty" binding:"required,IsAddress"`
	Amount  int    `json:"amount,omitempty" binding:"required,gt=0"`
}
