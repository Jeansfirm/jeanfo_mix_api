package middleware

// import (
// 	"net/http"

// 	"jeanfo_mix/internal/service"

// 	"github.com/gin-gonic/gin"
// )

// // AuthMiddleware 验证用户登录态
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token, err := c.Cookie("token")
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
// 			c.Abort()
// 			return
// 		}

// 		_, err = service.ParseToken(token)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录已失效"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
