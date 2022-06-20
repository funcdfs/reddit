package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/controller"
	"reddit/logger"
	middlewares "reddit/middleware"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 可以直接调用这一行，直接进入 release mode
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由

	v1 := r.Group("/api/v1")

	v1.POST("/sign_up", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("community", controller.CommunityHandler)
	}

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果当前的用户是登录的用户，判断请求的 header 中是否存在一个有效的 jwt token
		// 把认证的过程封装到 JWTAuthMiddleware 中间件中
		c.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 Not Found",
		})
	})
	return r
}
