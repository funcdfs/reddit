package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
定义程序中的响应内容，把响应的对象封装为固定的格式

code: 程序中的错误码
msg: xx 提示信息
data: 存放数据

对于前端来说处理的过程就会很清晰 code -> message -> data
*/

type ResponseData struct {
	Code ResCode
	Msg  any
	Data any
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMessage(c *gin.Context, code ResCode, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg() + msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
