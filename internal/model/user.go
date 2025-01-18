package model

import (
	"time"
)

// User 用户表模型
type User struct {
	ID            int    `gorm:"primaryKey"`              // 全局唯一用户ID
	Username      string `gorm:"type:varchar(50);unique"` // 用户名（普通用户登录使用）
	PasswordHash  string `gorm:"type:varchar(255)"`       // 加密后的密码
	Email         string `gorm:"type:varchar(100)"`       // 邮箱
	Phone         string `gorm:"type:varchar(20)"`        // 手机号
	Provider      string `gorm:"type:varchar(50)"`        // 第三方平台 (如"wechat", "github")，普通用户为空
	ProviderID    string `gorm:"type:varchar(100)"`       // 第三方平台用户ID，普通用户为空
	ProviderToken string `gorm:"type:varchar(255)"`       // 第三方平台token
	RegisterType  string `gorm:"varchar(16)"`             // 注册类型
	Status        int    `gorm:"type:tinyint;default:0"`  // 用户状态 0-正常 1-禁用
	Role          int    `gorm:"type:tinyint;default:1"`  // 用户角色 0-管理员 1-普通用户
	AccessToken   string `gorm:"type:varchar(255)"`       // 登录访问token
	RefreshToken  string `gorm:"type:varchar(255)"`       // 刷新token
	// TokenExpires   time.Time // token过期时间
	// RefreshExpires time.Time // 刷新token过期时间
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 更新时间
}
