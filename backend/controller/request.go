package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin     = errors.New("user not logged in")
	ErrorUserTokenToInt64 = errors.New("ErrorUserTokenToInt64")
)

const ContextUserId = "userID"

// GetCurrentUserId return userId int64 from contextUserId
func GetCurrentUserId(c *gin.Context) (userId uint64, err error) {
	uid, ok := c.Get(ContextUserId)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(uint64)
	if !ok {
		err = ErrorUserTokenToInt64
		return
	}

	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	SizeStr := c.Query("size")

	var (
		page int64 // 第几页 页数
		size int64 // 每页几条数据
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}

// localhost:8081/api/v1/posts/?size=1&page=2
