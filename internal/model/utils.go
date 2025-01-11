package model

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Demo{})
	db.AutoMigrate(&User{})
}
