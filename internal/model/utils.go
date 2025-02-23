package model

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Demo{})
	db.AutoMigrate(&User{})
	// blog app
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Comment{})
	// chat app
	db.AutoMigrate(&Conversation{})
	db.AutoMigrate(&Message{})
}
