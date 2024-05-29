package controller

import (
	"PaperSubmission/enum"
	"PaperSubmission/middleware"
	"PaperSubmission/model"
	"PaperSubmission/service"
	"PaperSubmission/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

func UserRegisterHandler(context *gin.Context) {
	email := context.PostForm("email")
	password := context.MustGet("password_sha256").(string) // 获取经过中间件层加密后的密码
	name := context.PostForm("name")
	organization := context.PostForm("organization")
	// 校验参数是否合法
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) || name == "" || organization == "" {
		// 参数不合法
		context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	user := model.User{
		Email:        email,
		Password:     password,
		Name:         name,
		Organization: organization,
	}
	if err := service.NewUserService().Add(user); err != nil {
		context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()))
}

func UserLoginHandler(context *gin.Context) {
	email := context.PostForm("email")
	password := context.MustGet("password_sha256").(string)                        // 获取经过中间件层加密后的密码
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // 邮箱正则表达式
	if !emailRegex.MatchString(email) {
		// 参数不合法
		context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	userID, name, err := service.NewUserService().VerifyPassword(email, password)
	if err != nil {
		context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	if userID == 0 { // 约定用户id为0则密码错误
		context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateFail), "密码错误"))
		return
	}
	// 设置cookie过期时间
	expires := time.Now().Add(7 * 24 * time.Hour)
	token, _ := middleware.ReleaseToken(userID)
	// 设置cookie, 将userID, username存储于cookie
	context.SetCookie("token", token, int(expires.Unix()), "/", "localhost:8080", true, false)
	context.SetCookie("username", name, int(expires.Unix()), "/", "localhost:8080", true, false)
	context.JSON(http.StatusOK, utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()))
}
