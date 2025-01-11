package test

import (
	"errors"
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

	model.MigrateDB(db)

	return db
}

func SetupTestDBV2() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed to connect to test datase: " + err.Error())
	}

	model.MigrateDB(db)

	return db, nil
}

func SetupTestRouter(db *gorm.DB) *gin.Engine {
	return router.SetupRouter(db)
}

func LoadTestConfig() {
	os.Setenv("APP_ENV", "test")
	config.LoadConfig()
}
