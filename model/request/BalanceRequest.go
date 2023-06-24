package request

type BalanceRequest struct {
	ERC20TokenAddress string `json:"ERC20TokenAddress,omitempty" binding:"required,IsAddress"`
	AccountAddress    string `json:"accountAddress,omitempty" binding:"required,IsAddress"`
}
