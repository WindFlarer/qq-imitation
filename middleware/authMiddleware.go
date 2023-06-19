package middleware

import (
	"net/http"

	"wechat-imitation/helper/login_register"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		// 解析JWT令牌
		uc, err := login_register.AnalyzeToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// 将UserID保存到上下文中
		c.Request.Header.Set("UserAccount", uc.UserAccount)

		// 继续处理请求
		c.Next()
	}
}
