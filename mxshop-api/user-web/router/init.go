package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
	middlewares "mxshop-api/user-web/middleware"
	"net/http"
)

func Init() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		})
	})

	engine.Use(middlewares.Cors())

	baseGroup := engine.Group("/u/v1/base")
	{
		baseGroup.POST("/login", api.PasswordLogin)
		baseGroup.GET("/captcha", api.GetCaptcha)
	}

	userGroup := engine.Group("/u/v1/user")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	}
	return engine
}
