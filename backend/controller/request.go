package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	middlewares "reddit/middleware"
)

var (
	ErrorUserNotLogin     = errors.New("user not logged in")
	ErrorUserTokenToInt64 = errors.New("ErrorUserTokenToInt64")
)

// GetCurrentUserID 返回 userID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.ContextUserId)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserTokenToInt64
		return
	}

	return
}
