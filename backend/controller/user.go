package controller

import (
	"errors"
	"reddit/dao/mysql"
	"reddit/logic"
	"reddit/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 参数校验:

	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数存在错误，记录日志，同时返回到前端
		zap.L().Error("signup ShouldBindJSON failed: ", zap.Error(err))
		ResponseError(c, CodeInvalidParameter)
		return
	}

	// 业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExists) {
			ResponseError(c, CodeUserExists)
			return
		}
		ResponseErrorWithMessage(c, CodeServerBusy, err.Error())
		return
	}
	// 返回响应
	ResponseSuccess(c, "sign_up successfully")
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数存在错误
		zap.L().Error("login ShouldBindJSON failed: ", zap.Error(err))
		ResponseError(c, CodeInvalidParameter)
		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("LoginHandler execute failed: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotFound)
			return
		}
		ResponseErrorWithMessage(c, CodeServerBusy, "controller.Login failed: "+err.Error())
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}
