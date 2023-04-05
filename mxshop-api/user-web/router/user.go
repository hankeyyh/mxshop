package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
)

func InitUserRouter(group *gin.RouterGroup) {
	g := group.Group("user")
	{
		g.GET("list", api.GetUserList)
	}
}
