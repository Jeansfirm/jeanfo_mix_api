package model

import (
	"time"
)

// User 用户表模型
type User struct {
	ID           string    `gorm:"primaryKey"`              // 全局唯一用户ID
	Username     string    `gorm:"type:varchar(50);unique"` // 用户名（普通用户登录使用）
	PasswordHash string    `gorm:"type:varchar(255)"`       // 加密后的密码
	Provider     string    `gorm:"type:varchar(50)"`        // 第三方平台 (如"wechat", "github")，普通用户为空
	ProviderID   string    `gorm:"type:varchar(100)"`       // 第三方平台用户ID，普通用户为空
	CreatedAt    time.Time `gorm:"autoCreateTime"`          // 创建时间
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`          // 更新时间
}
