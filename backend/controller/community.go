package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"reddit/logic"
)

// CommunityHandler implements the Handler interface for the Community router
// Response the community list to fronted
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed: ", zap.Error(err))
		ResponseError(c, CodeServerBusy) // don't return the server error to the frontend
		return
	}

	ResponseSuccess(c, data)
}
