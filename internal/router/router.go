package router

import (
	"jeanfo_mix/config"
	_ "jeanfo_mix/docs"
	"jeanfo_mix/internal/controller"
	"jeanfo_mix/internal/middleware"
	"jeanfo_mix/internal/service"
	chat_service "jeanfo_mix/internal/service/chat"
	user_service "jeanfo_mix/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	config := config.GetConfig()
	r := gin.Default()

	r.Use(middleware.RecoverMiddleWare())

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	demoService := service.DemoService{DB: db}
	demoController := controller.DemoController{Service: &demoService}
	userService := user_service.UserService{DB: db}
	userController := controller.UserController{Service: &userService}
	fileService := service.FileService{DB: db}
	fileController := controller.FileController{Service: &fileService}
	blogService := service.BlogService{DB: db}
	blogController := controller.BlogController{Service: &blogService}
	chatService := chat_service.ChatService{DB: db}
	chatController := controller.ChatController{Service: &chatService}

	// only demo
	r.GET("/api/demos/hello", demoController.HelloWorld)
	r.GET("/api/demos", demoController.GetDemos)
	r.POST("/api/demos", demoController.CreateDemo)
	r.DELETE("/api/demos/:id", demoController.DeleteDemo)

	// static files
	{
		// e.g. req_path: /static/upload/xx/1.jpg  => local_path: /to/upload_dir/xx/1.jpg
		r.StaticFS(config.Web.UploadDirStaticPath, http.Dir(config.Web.UploadDir))
	}

	// nologin apis
	noLoginApiGroup := r.Group("/api")
	{
		// nologin auth  ////////////////////////////////////////////////////////
		noLoginApiGroup.POST("/auth/register", userController.Register) // 用户注册
		noLoginApiGroup.POST("/auth/login", userController.Login)
	}

	// login apis ///////////////////////////////////////////////////////////////
	loginApiGroup := r.Group("/api")
	loginApiGroup.Use(middleware.AuthMiddleware()) // 需要登录态的接口

	// auth
	{
		loginApiGroup.POST("/auth/logout", userController.Logout)
		loginApiGroup.POST("/auth/change_passwd", userController.ChangePasswd)
	}

	// user
	{
		loginApiGroup.GET("/users/my", userController.GetUserMy)     // 获取我的用户信息
		loginApiGroup.POST("/users/my", userController.UpdateUserMy) // 更新我的用户信息
		loginApiGroup.GET("/users", userController.List)             // 获取用户列表
	}

	// file
	{
		loginApiGroup.POST("/file/upload", fileController.UploadFile)
		loginApiGroup.GET("/file/download", fileController.DownloadFile)
	}

	// blog
	{
		loginApiGroup.POST("/blog/articles", blogController.CreateArticle)
		loginApiGroup.GET("/blog/articles", blogController.ListArticle)
		loginApiGroup.GET("/blog/articles/my", blogController.ListArticleMy)

		loginApiGroup.POST("/blog/comments", blogController.CreateComment)
		loginApiGroup.GET("/blog/comments", blogController.ListComment)
	}

	// chat
	{
		chatGroup := loginApiGroup.Group("/chat")
		chatGroup.GET("/conversations", chatController.ListConversation)
		chatGroup.POST("/conversations", chatController.CreateConversion)
		chatGroup.GET("/messages", chatController.ListMessage)
	}

	return r
}
