package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"reddit/logic"
	"reddit/models"
)

func CreatePostHandler(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil { // validator --> binding tag
		zap.L().Debug("c.ShouldBindJSON(post) err", zap.Any("err", err))
		zap.L().Error("create post with invalid parm")
		ResponseErrorWithMessage(c, CodeInvalidParameter, err.Error())
		return
	}
	// 参数校验

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := GetCurrentUserId(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorId = uint64(userID)

	// 2、创建帖子
	err = logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccessWithData(c, nil)
}
