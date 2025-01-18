package model

import "time"

type Article struct {
	ID        int       `gorm:"primaryKey"`
	Ttile     string    `gorm:"type:varchar(256)"`
	Content   string    `gorm:"type:text"`
	PlainText string    `gorm:"type:text"`
	UserID    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 更新时间
}

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	Content   string    `gorm:"type:text"`
	PlainText string    `gorm:"type:text"`
	ArticleID int       `gorm:"type:int"`
	UserID    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 更新时间
}
