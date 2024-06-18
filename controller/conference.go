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
	"time"
)

type OneConferenceResponse struct {
	Conference
	response.Response
}

type ConferenceListResponse struct {
	List []Conference `json:"list"`
	response.Response
}

type Conference struct {
	Abbreviation     string    `json:"abbreviation"`
	CCfRanking       string    `json:"ccf_ranking"`
	FullName         string    `json:"full_name"`
	ID               int64     `json:"id"`
	Info             string    `json:"info"`
	Link             string    `json:"link"`
	MeetingDate      time.Time `json:"meeting_date"`
	MeetingVenue     string    `json:"meeting_venue"`
	MaterialDeadline time.Time `json:"material_deadline"`
	NotificationDate time.Time `json:"notification_date"`
	Sessions         int64     `json:"sessions"`
}

func GetConferenceHandle(context *gin.Context) {
	conferenceIDStr := context.Query("id")
	conferenceID, err := strconv.ParseInt(conferenceIDStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, OneConferenceResponse{Response: response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	conference := model.Conference{ID: conferenceID}
	conferenceModel, err := service.NewConferenceService().Get(conference)
	if err != nil {
		context.JSON(http.StatusInternalServerError, OneConferenceResponse{Response: response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	context.JSON(http.StatusOK, OneConferenceResponse{
		Conference: Conference{
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
		},
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func ConferenceListHandler(context *gin.Context) {
	pageStr := context.Query("page")
	pageSizeStr := context.Query("page_size")
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil { // 请求参数无法被解析
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	request := utils.ListQuery{Page: page, PageSize: pageSize}
	conferenceModelList, err := service.NewConferenceService().GetList(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var conferenceList []Conference
	for _, conferenceModel := range conferenceModelList {
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

func QueryConferenceHandle(context *gin.Context) {
	key := context.Query("key")
	conferenceModelList, err := model.NewConferenceModel().Query(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var conferenceList []Conference
	for _, conferenceModel := range conferenceModelList {
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
