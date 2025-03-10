package model

import (
	"jeanfo_mix/util"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID           uint      `gorm:"primaryKey"`
	MetaID       string    `gorm:"type:varchar(32);unique"`
	UserID       uint      `gorm:"type:int"`
	FileName     string    `gorm:"type:varchar(64)"`
	RelativePath string    `gorm:"type:varchar(128)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt    time.Time `gorm:"autoUpdateTime"` // 更新时间
}

func (f *File) Create(DB *gorm.DB) error {
	if f.MetaID == "" {
		f.MetaID = util.GenTimeBasedUUID(28)
	}
	err := DB.Create(f).Error

	return err
}
