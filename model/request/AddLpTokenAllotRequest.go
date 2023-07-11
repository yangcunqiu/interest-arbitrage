package request

type AddLpTokenAllotRequest struct {
	LpToken    string `json:"lpToken,omitempty" binding:"required,IsAddress"`
	AllocPoint int    `json:"allocPoint,omitempty" binding:"required"`
	IsUpdate   bool   `json:"isUpdate,omitempty"`
}
