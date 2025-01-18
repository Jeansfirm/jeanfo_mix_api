package router

import (
	_ "jeanfo_mix/docs"
	"jeanfo_mix/internal/controller"
	"jeanfo_mix/internal/middleware"
	"jeanfo_mix/internal/service"
	user_service "jeanfo_mix/internal/service/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.RecoverMiddleWare())

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	demoService := service.DemoService{DB: db}
	demoController := controller.DemoController{Service: &demoService}
	userService := user_service.UserService{DB: db}
	userController := controller.UserController{Service: &userService}

	// only demo
	r.GET("/api/demos/hello", demoController.HelloWorld)
	r.GET("/api/demos", demoController.GetDemos)
	r.POST("/api/demos", demoController.CreateDemo)
	r.DELETE("/api/demos/:id", demoController.DeleteDemo)

	// nologin auth
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", userController.Register) // 用户注册
		authGroup.POST("/login", userController.Login)
	}

	// login auth
	loginAuthGroup := r.Group("/api/auth")
	loginAuthGroup.Use(middleware.AuthMiddleware()) // 需要登录态的接口
	{
		loginAuthGroup.POST("/logout", userController.Logout)
		loginAuthGroup.POST("/change_passwd", userController.ChangePasswd)
	}

	// nologin apis ///////////////////////////////////////////////////////////
	// apiGroup := r.Group("/api")

	// login apis /////////////////////////////////////////////////////////////
	loginApiGroup := r.Group("/api")
	loginApiGroup.Use(middleware.AuthMiddleware()) // 需要登录态的接口

	// user
	{
		loginApiGroup.GET("/user", userController.Get)   // 获取用户信息
		loginApiGroup.GET("/users", userController.List) // 获取用户列表
	}

	// other
	{

	}

	return r
}
