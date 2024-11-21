package router

import (
	"jeanfo_mix/internal/controller"
	"jeanfo_mix/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	demoService := service.DemoService{DB: db}
	demoController := controller.DemoController{Service: &demoService}

	r.GET("/api/demos/hello", demoController.HelloWorld)
	r.GET("/api/demos", demoController.GetDemos)
	r.POST("/api/demos", demoController.CreateDemo)
	r.DELETE("/api/demos/:id", demoController.DeleteDemo)

	return r
}
