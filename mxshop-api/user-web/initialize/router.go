package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop-api/user-web/router"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	apiGroup := engine.Group("/u/v1")
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)
	zap.S().Info("InitRouter Suc")
	return engine
}
