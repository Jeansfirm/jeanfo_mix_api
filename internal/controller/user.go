package controller

// import (
// 	"net/http"

// 	"jeanfo_mix/internal/service"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// // GetUserInfoHandler 获取用户信息接口
// func GetUserInfoHandler(c *gin.Context) {
// 	token, _ := c.Cookie("token")
// 	claims, err := service.ParseToken(token)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "登录已失效"})
// 		return
// 	}

// 	userID := claims["user_id"].(string)
// 	db := c.MustGet("db").(*gorm.DB)
// 	user, err := service.GetUserByID(db, userID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id":    user.ID,
// 		"username":   user.Username,
// 		"provider":   user.Provider,
// 		"created_at": user.CreatedAt,
// 	})
// }
