package model

import (
	"time"
)

// User 用户表模型
type User struct {
	ID             string    `gorm:"primaryKey"`               // 全局唯一用户ID
	Username       string    `gorm:"type:varchar(50);unique"`  // 用户名（普通用户登录使用）
	PasswordHash   string    `gorm:"type:varchar(255)"`        // 加密后的密码
	Email          string    `gorm:"type:varchar(100);unique"` // 邮箱
	Phone          string    `gorm:"type:varchar(20);unique"`  // 手机号
	Provider       string    `gorm:"type:varchar(50)"`         // 第三方平台 (如"wechat", "github")，普通用户为空
	ProviderID     string    `gorm:"type:varchar(100)"`        // 第三方平台用户ID，普通用户为空
	ProviderToken  string    `gorm:"type:varchar(255)"`        // 第三方平台token
	Status         int       `gorm:"type:tinyint;default:1"`   // 用户状态 1-正常 2-禁用
	Role           int       `gorm:"type:tinyint;default:1"`   // 用户角色 1-普通用户 2-管理员
	ResetToken     string    `gorm:"type:varchar(100)"`        // 密码重置token
	ResetExpires   time.Time // 密码重置token过期时间
	AccessToken    string    `gorm:"type:varchar(255)"` // 登录访问token
	RefreshToken   string    `gorm:"type:varchar(255)"` // 刷新token
	TokenExpires   time.Time // token过期时间
	RefreshExpires time.Time // 刷新token过期时间
	CreatedAt      time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt      time.Time `gorm:"autoUpdateTime"` // 更新时间
}
