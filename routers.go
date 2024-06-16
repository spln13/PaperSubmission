package main

import (
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
	server.GET("/register/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.html", "")
	})
	server.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.html", "")
	})
	server.GET("/conferences/:page/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "conference.html", "")
	})
	server.GET("/journals/:page/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "journal.html", "")
	})
	server.GET("/special_issues/:page/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "special_issue.html", "")
	})
	server.GET("/conference/:conference_id/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "one_conference.html", "")
	})
	server.GET("/journal/:journal_id/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "one_journal.html", "")
	})
	server.GET("/special_issue/:special_issue_id/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "one_special_issue.html", "")
	})
	server.GET("/followed_conferences/:page/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "followed_conferences.html", "")
	})
	server.GET("followed_journals/:page/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "followed_journals.html", "")
	})
	server.GET("/logout/", func(context *gin.Context) {
		context.SetCookie("token", "", -1, "/", "127.0.0.1:8080", true, false)
		context.SetCookie("username", "", -1, "/", "127.0.0.1:8080", true, false)
		context.Redirect(http.StatusFound, "/")
	})
	// api接口
	userGroup := server.Group("/api/user")
	{
		userGroup.POST("/register/", middleware.PasswordEncryptionMiddleware(), controller.UserRegisterHandler)
		userGroup.POST("/login/", middleware.PasswordEncryptionMiddleware(), controller.UserLoginHandler)
		userGroup.GET("/following_journals/", middleware.UserJWTMiddleware(), controller.GetUserFollowingJournalListHandler)
		userGroup.GET("/following_conferences/", middleware.UserJWTMiddleware(), controller.GetUserFollowingConferenceListHandler)
	}

	homeGroup := server.Group("/api/home")
	{
		homeGroup.GET("/information/", controller.HomeInformationHandle)
		homeGroup.GET("/list/", controller.HomeListHandler)
	}
	journalGroup := server.Group("/api/journal")
	{
		journalGroup.GET("/get/", middleware.UserJWTMiddleware(), controller.GetJournalHandle)
		journalGroup.GET("/list/", middleware.UserJWTMiddleware(), controller.JournalListHandler)
		journalGroup.POST("/follow/", middleware.UserJWTMiddleware(), controller.FollowJournalHandler)
		journalGroup.POST("/unfollow/", middleware.UserJWTMiddleware(), controller.UnfollowJournalHandler)
		journalGroup.GET("/followed_users/", middleware.UserJWTMiddleware(), controller.GetJournalFollowedUserListHandler)
	}
	conferenceGroup := server.Group("/api/conference")
	{
		conferenceGroup.GET("/get/", middleware.UserJWTMiddleware(), controller.GetConferenceHandle)
		conferenceGroup.GET("/list/", middleware.UserJWTMiddleware(), controller.ConferenceListHandler)
		conferenceGroup.POST("/follow/", middleware.UserJWTMiddleware(), controller.FollowConferenceHandler)
		conferenceGroup.POST("/unfollow/", middleware.UserJWTMiddleware(), controller.UnfollowConferenceHandler)
		conferenceGroup.GET("/followed_users/", middleware.UserJWTMiddleware(), controller.GetConferenceFollowedUserListHandler)
	}
	specialIssueGroup := server.Group("/api/special_issue")
	{
		specialIssueGroup.GET("/get/", middleware.UserJWTMiddleware(), controller.GetSpecialIssueHandle)
		specialIssueGroup.GET("/list/", middleware.UserJWTMiddleware(), controller.SpecialIssueListHandler)
	}
	return server
}
