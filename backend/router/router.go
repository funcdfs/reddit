package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"reddit/controller"
	"reddit/logger"
	"reddit/middleware"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		// call this line directly to enter release mode
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// use logger middleware

	// register service routing
	v1 := r.Group("/api/v1")
	v1.POST("/sign_up", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	// v1 group use jwtAuthMiddleware
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("community", controller.CommunityHandler)
	}

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		// if the current user is a sign in user
		// determine whether there is a valid jwt token in the request header
		// if jwt token is valid then send the pong
		c.String(http.StatusOK, "pong")
	})

	// change noRoute to 404 page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 Not Found",
		})
	})
	return r
}
