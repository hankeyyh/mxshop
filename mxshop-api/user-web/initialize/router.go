package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	middlewares "mxshop-api/user-web/middleware"
	"mxshop-api/user-web/router"
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
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)
	zap.S().Info("InitRouter Suc")
	return engine
}
