package model

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"size:255"`
	Summary   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 更新时间
	DeleteAt  gorm.DeletedAt
	Messages  []Message
}

type Message struct {
	ID             uint      `gorm:"primaryKey"`
	ConversationID uint      `gorm:"not null"`
	Role           string    `gorm:"size:20;not null"`
	ContentType    string    `gorm:"size:20;not null"`
	Content        string    `gorm:"type:text;not null"`
	Sequence       int       `gorm:"not null"`
	CreateAt       time.Time `gorm:"type:datetime(3)"`
}
