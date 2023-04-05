package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/router"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	apiGroup := engine.Group("/u//v1")
	router.InitUserRouter(apiGroup)
	return engine
}
