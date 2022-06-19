package middlewares

import (
	"github.com/gin-gonic/gin"
	"reddit/controller"
	"reddit/pkg/jwt"
	"strings"
)

const ContextUserId = "userID"

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带 Token 有三种方式 1：放在请求头，2：放在请求体，3：放在 URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
		// Authorization: Bearer xxx.xxx.xxx
		// 这里的具体实现方式要依据实际的业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// 按照空格进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString, 我们使用之前定义好的解析 JWT 的函数来解析他
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return
		}

		// 将当前请求的 userID 信息保存到请求的上下文 c 上面
		c.Set(ContextUserId, mc.UserID)
		c.Next() // 后续的处理函数可以用过 c.Get(ContextUserId) 来获取当前请求的用户信息

		// 得到 userID 就可以查询所有的数据了
	}
}
