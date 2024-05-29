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
	userGroup := server.Group("/api/user")
	{
		userGroup.POST("/register/", middleware.PasswordEncryptionMiddleware(), controller.UserRegisterHandler)
		userGroup.POST("/login/", middleware.PasswordEncryptionMiddleware(), controller.UserLoginHandler)
	}

	homeGroup := server.Group("/api/home")
	{
		homeGroup.GET("/information/", controller.HomeInformationHandle)
	}
	journalGroup := server.Group("/api/journal")
	{
		journalGroup.GET("/get/", middleware.UserJWTMiddleware(), controller.GetJournalHandle)
	}
	conferenceGroup := server.Group("/api/conference")
	{
		conferenceGroup.GET("/get/", middleware.UserJWTMiddleware(), controller.GetConferenceHandle)
	}
	return server
}
