package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"reddit/dao/mysql"
	"reddit/logic"
	"reddit/models"
)

func SignUpHandler(c *gin.Context) {
	// 参数校验:
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数存在错误
		zap.L().Error("signup ShouldBindJSON failed: ", zap.Error(err))
		ResponseError(c, CodeInvalidParameter)
		return
	}

	// 手动对请求的参数进行详细的业务规则校验，不要相信前端发送过来的数据
	// if len(p.UserName) == 0 || len(p.Password) == 0 || p.RePassword != p.Password {
	// 	c.JSON(http.StatusOK, gin.H{"error": "请求参数格式有误"})
	// }
	// 对于参数校验可以使用 validator 库进行校验

	// 业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExists) {
			ResponseError(c, CodeUserExists)
			return
		}
		ResponseErrorWithMessage(c, CodeServerBusy, "code server_busy")
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
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotFound)
			return
		}
		ResponseErrorWithMessage(c, CodeInvalidParameter, "logic.Login failed"+err.Error())
		return
	}
	// 返回响应
	ResponseSuccess(c, "login successful")
}
