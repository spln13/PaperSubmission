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

type OneJournalResponse struct {
	Abbreviation string    `json:"abbreviation"`
	CCFRanking   string    `json:"ccf_ranking"`
	Deadline     time.Time `json:"deadline"`
	Description  string    `json:"description"`
	FullName     string    `json:"full_name"`
	ID           int64     `json:"id"`
	ImpactFactor float64   `json:"impact_factor"`
	ISSN         string    `json:"issn"`
	Publisher    string    `json:"publisher"`
	utils.Response
}

func GetJournalHandle(context *gin.Context) {
	journalIDStr := context.Query("id")
	// 将journalIDStr转换为int64
	journalID, err := strconv.ParseInt(journalIDStr, 10, 64)
	if err != nil {
		// journalID无法被转换为int64, 返回错误
		context.JSON(http.StatusOK, OneJournalResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	journal := model.Journal{ID: journalID}
	journalModel, err := service.NewJournalService().Get(journal)
	if err != nil {
		context.JSON(http.StatusOK, OneJournalResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String())})
		return
	}
	// 返回信息
	context.JSON(http.StatusOK, OneJournalResponse{
		Abbreviation: journalModel.Abbreviation,
		CCFRanking:   journalModel.CCFRanking,
		Deadline:     journalModel.Deadline,
		Description:  journalModel.Description,
		FullName:     journalModel.FullName,
		ID:           journalModel.ID,
		ImpactFactor: journalModel.ImpactFactor,
		ISSN:         journalModel.ISSN,
		Publisher:    journalModel.Publisher,
		Response:     utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
