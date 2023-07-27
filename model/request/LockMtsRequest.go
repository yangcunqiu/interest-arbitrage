package request

type LockMtsRequest struct {
	UserAddress string `json:"userAddress,omitempty" binding:"required,IsAddress"`
	ChainId     int    `json:"chainId,omitempty" binding:"required"`
	Amount      string `json:"amount,omitempty" binding:"required,gt=0"`
}
