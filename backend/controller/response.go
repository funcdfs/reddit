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
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
