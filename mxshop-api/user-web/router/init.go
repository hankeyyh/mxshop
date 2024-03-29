package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/api"
	middlewares "github.com/hankeyyh/mxshop/mxshop-api/user-web/middleware"
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
		baseGroup.GET("/captcha", api.GetCaptcha)
	}

	userGroup := engine.Group("/u/v1/user")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		userGroup.POST("/login", api.PasswordLogin)
		userGroup.POST("/register", api.Register)
	}
	return engine
}
