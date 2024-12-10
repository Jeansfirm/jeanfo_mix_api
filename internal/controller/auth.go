package controller

// import (
// 	"net/http"

// 	"jeanfo_mix/internal/service"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// // 用户注册接口
// func RegisterHandler(c *gin.Context) {
// 	var req struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
// 		return
// 	}

// 	db := c.MustGet("db").(*gorm.DB)
// 	user, err := service.Register(db, req.Username, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "message": "注册成功"})
// }

// // 用户登录接口
// func LoginHandler(c *gin.Context) {
// 	var req struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
// 		return
// 	}

// 	db := c.MustGet("db").(*gorm.DB)
// 	user, token, err := service.Login(db, req.Username, req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 设置Cookie
// 	c.SetCookie("token", token, 72*3600, "/", "", false, true)
// 	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID})
// }

// // 第三方登录接口
// func ThirdPartyLoginHandler(c *gin.Context) {
// 	var req struct {
// 		Provider   string `json:"provider"`
// 		ProviderID string `json:"provider_id"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
// 		return
// 	}

// 	db := c.MustGet("db").(*gorm.DB)
// 	user, token, err := service.ThirdPartyLogin(db, req.Provider, req.ProviderID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID})
// }
