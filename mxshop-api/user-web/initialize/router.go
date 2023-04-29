package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/api"
	middlewares "mxshop-api/user-web/middleware"
	"net/http"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	// 健康检查
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域
	engine.Use(middlewares.Cors())

	apiGroup := engine.Group("/u/v1")
	InitUserRouter(apiGroup)
	InitBaseRouter(apiGroup)
	zap.S().Info("InitRouter Suc")
	return engine
}

func InitBaseRouter(group *gin.RouterGroup) {
	g := group.Group("base")
	{
		g.POST("login", api.PasswordLogin)
		g.GET("captcha", api.GetCaptcha)
	}
}

func InitUserRouter(group *gin.RouterGroup) {
	g := group.Group("user")
	{
		g.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	}
}
