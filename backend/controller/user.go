package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"reddit/logic"
	"reddit/models"
)

func SignUpHandler(c *gin.Context) {
	// 参数校验:
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数存在错误
		zap.L().Error("ShouldBindJSON failed: ", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 手动对请求的参数进行详细的业务规则校验，不要相信前端发送过来的数据
	// if len(p.UserName) == 0 || len(p.Password) == 0 || p.RePassword != p.Password {
	// 	c.JSON(http.StatusOK, gin.H{"error": "请求参数格式有误"})
	// }
	// 对于参数校验可以使用 validator 库进行校验

	// 业务处理
	if err := logic.SignUp(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, "ok")
}
