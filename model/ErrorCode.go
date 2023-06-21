package model

type ErrorCode struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
