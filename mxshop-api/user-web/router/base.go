package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
)

func InitBaseRouter(group *gin.RouterGroup) {
	g := group.Group("base")
	{
		g.POST("login", api.PasswordLogin)
	}
}
