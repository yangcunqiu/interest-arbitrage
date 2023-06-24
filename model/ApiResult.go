package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResult struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data any, message ...string) {
	msg := "success"
	for _, str := range message {
		msg += " " + str
	}

	c.JSON(http.StatusOK, ApiResult{
		Code:    200,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, errorCode ErrorCode, customMessage ...string) {
	errorMessage := errorCode.Message
	for _, str := range customMessage {
		errorMessage += " " + str
	}
	c.JSON(http.StatusOK, ApiResult{
		Code:    errorCode.Code,
		Message: errorMessage,
	})
}
