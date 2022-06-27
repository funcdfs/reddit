package controller

import (
	"strconv"

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

	ResponseSuccessWithData(c, data)
}

// CommunityDetailHandler
func CommunityDetailHandler(c *gin.Context) {
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParameter)
		return
	}

	communityList, err := logic.GetCommunityDetailByID(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID() failed", zap.Error(err))
		ResponseErrorWithMessage(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccessWithData(c, communityList)
}
