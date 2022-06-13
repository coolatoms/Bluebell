package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 错误响应
func ResponseError(ctx *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMsg 自定义错误响应
func ResponseErrorWithMsg(ctx *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

// ResponseSuccess 相应正确
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}
