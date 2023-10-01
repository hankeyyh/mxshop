package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/constant"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*CustomClaims)
		if currentUser.AuthorityId != constant.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
