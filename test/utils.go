package test

import (
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test datase: %v", err)
	}

	if err := db.AutoMigrate(&model.Demo{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

func SetupTestRouter(db *gorm.DB) *gin.Engine {
	return router.SetupRouter(db)
}

func LoadTestConfig() {
	os.Setenv("APP_ENV", "test")
	config.LoadConfig()
}
