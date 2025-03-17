package model

import (
	"fmt"
	"jeanfo_mix/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var initDBLock sync.Mutex

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Demo{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&File{})
	// blog app
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Comment{})
	// chat app
	db.AutoMigrate(&Conversation{})
	db.AutoMigrate(&Message{})
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	initDBLock.Lock()
	defer initDBLock.Unlock()
	if db == nil {
		cfg := config.GetConfig()
		dbConfig := cfg.Database

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
		ndb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database: %v", err))
		}

		MigrateDB(ndb)
		db = ndb
	}

	return db
}
