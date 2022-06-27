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

// GetCurrentUserId return userId int64 from contextUserId
func GetCurrentUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserId)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserTokenToInt64
		return
	}

	return
}
