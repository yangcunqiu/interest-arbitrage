package model

var (
	CommonError    = ErrorCode{10000, "fail: "}
	ParamBindError = ErrorCode{10001, "参数绑定错误"}

	NotAddress = ErrorCode{20000, "不是一个地址"}
)
