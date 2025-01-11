package auth_service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// 生成随机用户名，格式类似 用户-{短uuid}
func GenerateUserName() string {
	return "用户-" + uuid.New().String()[24:]
}
