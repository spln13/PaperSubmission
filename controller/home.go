package controller

import (
	"PaperSubmission/cache"
	"PaperSubmission/enum"
	"PaperSubmission/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeCounterResponse struct {
	Conferences int64 `json:"conferences"`
	Journals    int64 `json:"journals"`
	Users       int64 `json:"users"`
	PageViews   int64 `json:"page_views"`
	response.Response
}

type HomeListResponse struct {
	response.Response
	ConferenceList   []Conference            `json:"conference_list"`
	JournalList      []Journal               `json:"journal_list"`
	SpecialIssueList []response.SpecialIssue `json:"special_issue_list"`
}

func HomeInformationHandle(context *gin.Context) {
	// 先增加页面浏览量
	_ = cache.NewHomePageCache().IncreasePageView()
	conferenceNum, err := cache.NewHomePageCache().GetConferenceNum() // 从cache获取会议数量
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	journalNum, err := cache.NewHomePageCache().GetJournalNum()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	userNum, err := cache.NewHomePageCache().GetUserNum()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	pageViews, err := cache.NewHomePageCache().GetPageView()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	// 操作成功，返回参数
	context.JSON(http.StatusOK, HomeCounterResponse{
		Conferences: conferenceNum,
		Journals:    journalNum,
		Users:       userNum,
		PageViews:   pageViews,
		Response:    response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func HomeListHandler(context *gin.Context) {
	conferenceModelList, err := cache.NewHomePageCache().GetConferenceList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	journalModelList, err := cache.NewHomePageCache().GetJournalList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	specialIssueList, err := cache.NewHomePageCache().GetSpecialIssueList()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var conferenceList []Conference
	var journalList []Journal
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
	for _, journalModel := range journalModelList {
		journal := Journal{
			Abbreviation: journalModel.Abbreviation,
			CCFRanking:   journalModel.CCFRanking,
			Deadline:     journalModel.Deadline,
			Description:  journalModel.Description,
			FullName:     journalModel.FullName,
			ID:           journalModel.ID,
			ImpactFactor: journalModel.ImpactFactor,
			ISSN:         journalModel.ISSN,
			Publisher:    journalModel.Publisher,
		}
		journalList = append(journalList, journal)
	}
	context.JSON(http.StatusOK, HomeListResponse{
		Response:         response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
		ConferenceList:   conferenceList,
		JournalList:      journalList,
		SpecialIssueList: specialIssueList,
	})
}
