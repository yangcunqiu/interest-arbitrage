package request

type TokenApprove struct {
	SpenderAddress    string `json:"spenderAddress,omitempty" binding:"required,IsAddress"`
	ERC20TokenAddress string `json:"ERC20TokenAddress,omitempty" binding:"required,IsAddress"`
	Amount            uint   `json:"amount,omitempty" binding:"required,gt=0"`
}
