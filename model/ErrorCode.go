package model

type ErrorCode struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func ErrorCodeOf(code int, message string) ErrorCode {
	return ErrorCode{code, message}
}
