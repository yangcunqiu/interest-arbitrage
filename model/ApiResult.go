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

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResult{
		Code:    200,
		Message: "success",
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
