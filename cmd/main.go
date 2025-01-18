package main

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"
	"jeanfo_mix/util/log_util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	@title			JEANFO_MIX_API
//	@version		1.0
//	@description	This is a WEB server for JEANFO_MIX_API.
//	@termsOfService	jeanfo.cn

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @contact.name	Jeanfo Peng
// @contact.url	http://jeanfo.cn
// @contact.email	jeanf@qq.com
func main() {
	config.LoadConfig()
	webConfig := config.AppConfig.Web
	dbConfig := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log_util.Fatal("Failed to connect to database: %v", err)
	}

	model.MigrateDB(db)

	listen_on := fmt.Sprintf("%s:%d", webConfig.Host, webConfig.Port)
	r := router.SetupRouter(db)
	r.Run(listen_on)
}
