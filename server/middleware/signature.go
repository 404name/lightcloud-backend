package middleware

import "github.com/gin-gonic/gin"

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验签

		c.Next()

		// 加签
	}
}
