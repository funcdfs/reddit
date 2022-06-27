package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"reddit/dao/mysql"
	"reddit/logic"
	"reddit/models"
)

// SignUpHandler implements sign up route handler
// check the parameters is valid
// use logic.SignUp to sign up the user
// if the user sign up successfully then return success to the frontend
func SignUpHandler(c *gin.Context) {
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signup ShouldBindJSON failed: ", zap.Error(err))
		ResponseError(c, CodeInvalidParameter)
		return
	}
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExists) {
			ResponseError(c, CodeUserExists)
			return
		}
		ResponseErrorWithMessage(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccessWithData(c, "sign_up successfully")
}

// LoginHandler implements login router handler
// check the parameters is valid
// use logic.Login to sign in the user that parameters specified
// is successfully sign in, then return the jwt token to frontend
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login ShouldBindJSON failed: ", zap.Error(err))
		ResponseError(c, CodeInvalidParameter)
		return
	}
	/*

		post format:

		{
			"username": "konng",
			"password": "1"
		}

	*/
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("LoginHandler execute failed: ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotFound)
			return
		}
		ResponseErrorWithMessage(c, CodeServerBusy, "logic.Login failed: "+err.Error())
		return
	}
	// response token to frontend
	ResponseSuccessWithData(c, token)
}
