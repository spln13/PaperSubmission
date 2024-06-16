package controller

import (
	"PaperSubmission/enum"
	"PaperSubmission/model"
	"PaperSubmission/response"
	"PaperSubmission/service"
	"PaperSubmission/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FollowConferenceHandler(context *gin.Context) {
	userID, _ := context.MustGet("userID").(int64)
	conferenceIDStr := context.Query("conference_id")
	conferenceID, err := strconv.ParseInt(conferenceIDStr, 10, 64)
	if err != nil { // 请求参数中journalID不合法
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	followConference := model.FollowConference{UserID: userID, ConferenceID: conferenceID}
	exist, err := service.NewFollowConferenceService().Exists(followConference)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	if exist == true {
		context.JSON(http.StatusOK, response.NewCommonResponse(int(enum.OperateFail), "已关注"))
		return
	}
	if err := service.NewFollowConferenceService().Add(followConference); err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	context.JSON(http.StatusOK, response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()))
}

func GetUserFollowingConferenceListHandler(context *gin.Context) {
	userID, _ := context.MustGet("userID").(int64)
	pageStr := context.Query("page")
	pageSizeStr := context.Query("page_size")
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil { // 请求参数无法被解析
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	request := utils.ListQuery{Page: page, PageSize: pageSize}
	conferences, err := service.NewFollowConferenceService().GetConferenceList(userID, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var conferenceList []Conference
	for _, conferenceModel := range conferences {
		conference := Conference{
			Abbreviation:     conferenceModel.Abbreviation,
			CCfRanking:       conferenceModel.CCFRanking,
			FullName:         conferenceModel.FullName,
			ID:               conferenceModel.ID,
			Info:             conferenceModel.Info,
			Link:             conferenceModel.Link,
			MeetingDate:      conferenceModel.MeetingDate,
			MeetingVenue:     conferenceModel.MeetingVenue,
			MaterialDeadline: conferenceModel.MaterialDeadline,
			NotificationDate: conferenceModel.NotificationDate,
			Sessions:         conferenceModel.Sessions,
		}
		conferenceList = append(conferenceList, conference)
	}
	context.JSON(http.StatusOK, ConferenceListResponse{
		List:     conferenceList,
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func GetConferenceFollowedUserListHandler(context *gin.Context) {
	conferenceIDStr := context.Query("conference_id")
	pageStr := context.Query("page")
	pageSizeStr := context.Query("page_size")
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil { // 请求参数无法被解析
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	request := utils.ListQuery{Page: page, PageSize: pageSize}
	conferenceID, err := strconv.ParseInt(conferenceIDStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	users, err := service.NewFollowConferenceService().GetUserList(conferenceID, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var userList []User
	for _, userModel := range users {
		user := User{
			ID:           userModel.ID,
			Name:         userModel.Name,
			Organization: userModel.Organization,
		}
		userList = append(userList, user)
	}
	context.JSON(http.StatusOK, UserListResponse{
		List:     userList,
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
