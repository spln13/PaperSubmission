package main

import (
	"PaperSubmission/cache"
	"PaperSubmission/controller"
	"PaperSubmission/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer() *gin.Engine {
	server := gin.Default()             // 初始化gin服务器
	server.Static("static", "./static") // 指定静态文件path
	server.LoadHTMLGlob("template/*")   // 指定HTML文件path
	// 返回HTML页面
	server.GET("/login/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", "")
	})
	server.GET("/", func(context *gin.Context) {
		// 增加页面浏览量
		_ = cache.NewHomePageCache().IncreasePageView()
		context.HTML(http.StatusOK, "home.html", "")
	})

	// api接口
	studentGroup := server.Group("/api/student")
	{
		studentGroup.POST("/register/", middleware.PasswordEncryptionMiddleware(), controller.UserRegisterHandler)
		studentGroup.POST("/login/", middleware.PasswordEncryptionMiddleware(), controller.UserLoginHandler)
	}

	homeGroup := server.Group("/api/home")
	{
		homeGroup.GET("/information/", controller.HomeInformationHandle)
	}
	journalGroup := server.Group("/api/journal")
	{
		journalGroup.GET("/get/", controller.GetJournalHandle)
	}
	conferenceGroup := server.Group("/api/conference")
	{
		conferenceGroup.GET("/get/", controller.GetConferenceHandle)
	}
	return server
}
