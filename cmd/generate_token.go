package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"jeanfo_mix/internal/model"
// 	"jeanfo_mix/internal/service"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // GenerateTokenForUser 生成用户Token
// func GenerateTokenForUser(db *gorm.DB, username string) {
// 	var user model.User
// 	err := db.Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		// 如果用户不存在，创建用户
// 		user = model.User{
// 			ID:           uuid.New().String(),
// 			Username:     username,
// 			PasswordHash: "", // 空密码，因为是命令行创建
// 		}
// 		if err := db.Create(&user).Error; err != nil {
// 			log.Fatalf("创建用户失败: %v", err)
// 		}
// 		fmt.Printf("用户 %s 不存在，已创建。\n", username)
// 	}

// 	token, err := service.GenerateToken(user.ID)
// 	if err != nil {
// 		log.Fatalf("生成Token失败: %v", err)
// 	}

// 	fmt.Printf("用户 %s 的Token: %s\n", username, token)
// 	os.Exit(0)
// }
