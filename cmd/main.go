package main

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()
	webConfig := config.AppConfig.Web
	dbConfig := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	model.MigrateDB(db)

	listen_on := fmt.Sprintf("%s:%d", webConfig.Host, webConfig.Port)
	r := router.SetupRouter(db)
	r.Run(listen_on)
}
