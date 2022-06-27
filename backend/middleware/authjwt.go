package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"reddit/controller"
	"reddit/pkg/jwt"
)

const ContextUserIDKey = "userID"

// JWTAuthMiddleware 基于 JWT 的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErrorWithMessage(c, controller.CodeNeedAuth, "need auth token")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseErrorWithMessage(c, controller.CodeInvalidToken, "token format is error")
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			fmt.Println(err)
			controller.ResponseErrorWithMessage(c, controller.CodeNeedLogin, "token parse error")
			c.Abort()
			return
		}
		// store userid into context
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next() // c.Get(ContextUserIDKey)
	}
}
