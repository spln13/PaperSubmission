package controller

import (
	"PaperSubmission/enum"
	"PaperSubmission/model"
	"PaperSubmission/service"
	"PaperSubmission/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type OneConferenceResponse struct {
	Conference
	utils.Response
}

type Conference struct {
	Abbreviation     string    `json:"abbreviation"`
	CCfRanking       string    `json:"ccf_ranking"`
	FullName         string    `json:"full_name"`
	ID               int64     `json:"id"`
	Info             string    `json:"info"`
	Link             string    `json:"link"`
	Location         string    `json:"location"`
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
		context.JSON(http.StatusBadRequest, OneConferenceResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	conference := model.Conference{ID: conferenceID}
	conferenceModel, err := service.NewConferenceService().Get(conference)
	if err != nil {
		context.JSON(http.StatusInternalServerError, OneConferenceResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
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
			Location:         conferenceModel.Location,
			MeetingDate:      conferenceModel.MeetingDate,
			MeetingVenue:     conferenceModel.MeetingVenue,
			MaterialDeadline: conferenceModel.MaterialDeadline,
			NotificationDate: conferenceModel.NotificationDate,
			Sessions:         conferenceModel.Sessions,
		},
		Response: utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
