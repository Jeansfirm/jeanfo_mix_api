package model

type Demo struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"type:varchar(255)"`
	Content string `gorm:"type:text"`
}
