package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	middlewares "mxshop-api/user-web/middleware"
	"mxshop-api/user-web/router"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	// 配置跨域
	engine.Use(middlewares.Cors())

	apiGroup := engine.Group("/u/v1")
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)
	zap.S().Info("InitRouter Suc")
	return engine
}
