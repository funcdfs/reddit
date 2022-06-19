package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin     = errors.New("user not logged in")
	ErrorUserTokenToInt64 = errors.New("ErrorUserTokenToInt64")
)

const ContextUserId = "userID"

// GetCurrentUserID 返回 userID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserId)
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
