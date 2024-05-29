package controller

import (
	"PaperSubmission/cache"
	"PaperSubmission/enum"
	"PaperSubmission/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeResponse struct {
	Conferences    int64 `json:"conferences"`
	Journals       int64 `json:"journals"`
	Users          int64 `json:"users"`
	PageViews      int64 `json:"page_views"`
	commonResponse utils.Response
}

func HomeInformationHandle(context *gin.Context) {
	conferenceNum, err := cache.NewHomePageCache().GetConferenceNum() // 从cache获取会议数量
	if err != nil {
		context.JSON(http.StatusOK, HomeResponse{commonResponse: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	journalNum, err := cache.NewHomePageCache().GetJournalNum()
	if err != nil {
		context.JSON(http.StatusOK, HomeResponse{commonResponse: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	userNum, err := cache.NewHomePageCache().GetUserNum()
	if err != nil {
		context.JSON(http.StatusOK, HomeResponse{commonResponse: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	pageViews, err := cache.NewHomePageCache().GetPageView()
	if err != nil {
		context.JSON(http.StatusOK, HomeResponse{commonResponse: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	// 操作成功，返回参数
	context.JSON(http.StatusOK, HomeResponse{
		Conferences:    conferenceNum,
		Journals:       journalNum,
		Users:          userNum,
		PageViews:      pageViews,
		commonResponse: utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
