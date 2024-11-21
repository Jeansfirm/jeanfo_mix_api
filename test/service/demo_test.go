package service_test

import (
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/service"
	"jeanfo_mix/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemoService_CreateDemo(t *testing.T) {
	db := test.SetupTestDB(t)
	demoService := service.DemoService{DB: db}

	demo, err := demoService.CreateDemo("test title", "test content")
	assert.NoError(t, err)
	assert.NotNil(t, demo)
	assert.Equal(t, "test title", demo.Title)

	var storedDemo model.Demo
	db.First(&storedDemo, demo.ID)
	assert.Equal(t, "test content", storedDemo.Content)
}

func TestDemoService_DeleteDemo(t *testing.T) {
	db := test.SetupTestDB(t)
	demoService := service.DemoService{DB: db}

	demo := &model.Demo{Title: "test", Content: "content"}
	db.Create(demo)

	err := demoService.DeleteDemo(demo.ID)
	assert.NoError(t, err)

	err = demoService.DeleteDemo(999)
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
}
