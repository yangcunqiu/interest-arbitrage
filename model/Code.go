package model

var (
	commonCode            = 10000
	CommonError           = ErrorCode{getCommonCode(), "fail: "}
	ParamBindError        = ErrorCode{getCommonCode(), "参数绑定错误"}
	FunctionInvokeTimeOut = ErrorCode{getCommonCode(), "函数调用超时"}

	contractCode                = 20000
	AddressNotBlankOrNotAddress = ErrorCode{getContractCode(), "地址不能为空或地址格式不正确"}
	GetContractError            = ErrorCode{getContractCode(), "获取合约失败:"}
	CallContractError           = ErrorCode{getContractCode(), "合约调用失败:"}
	BuildTransactOptsError      = ErrorCode{getContractCode(), "构建交易参数失败:"}
	WaitTransactionPackError    = ErrorCode{getContractCode(), "等待交易打包失败:"}
	WaitTransactionPackTimeOut  = ErrorCode{getContractCode(), "等待交易打包超时:"}
)

func getCommonCode() int {
	commonCode++
	return commonCode
}

func getContractCode() int {
	contractCode++
	return contractCode
}
