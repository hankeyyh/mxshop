package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
	middlewares "mxshop-api/user-web/middleware"
)

func InitUserRouter(group *gin.RouterGroup) {
	g := group.Group("user")
	{
		g.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
	}
}
