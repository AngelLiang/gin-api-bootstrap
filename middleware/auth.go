package middleware

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	config "gin_api_bootstrap/config"

	"gin_api_bootstrap/util"
)

// authMiddleware中间件会检查每个请求的身份验证信息是否正确
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.Request.Header.Get("Authorization")
		if (apikey != config.Config.ApiKey) {
			resp := util.MakeResponse(1, "认证失败", gin.H{})
			c.AbortWithStatusJSON(401, resp)
			return
		}
		// 继续处理请求
		c.Next()
	}
}
