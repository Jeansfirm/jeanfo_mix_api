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

	// auth
	// authGroup := r.Group("/api/auth")
	// {
	// 	authGroup.POST("/register", controller.RegisterHandler)           // 用户注册
	// 	authGroup.POST("/login", controller.LoginHandler)                 // 用户登录
	// 	authGroup.POST("/third_party", controller.ThirdPartyLoginHandler) // 第三方登录
	// }

	// loginApiGroup := r.Group("/api")
	// loginApiGroup.Use(middleware.AuthMiddleware()) // 需要登录态的接口
	// {
	// 	loginApiGroup.GET("/user/info", controller.GetUserInfoHandler) // 获取用户信息
	// }

	return r
}
