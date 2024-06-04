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
	Journal
	utils.Response
}
type Journal struct {
	Abbreviation string    `json:"abbreviation"`
	CCFRanking   string    `json:"ccf_ranking"`
	Deadline     time.Time `json:"deadline"`
	Description  string    `json:"description"`
	FullName     string    `json:"full_name"`
	ID           int64     `json:"id"`
	ImpactFactor float64   `json:"impact_factor"`
	ISSN         string    `json:"issn"`
	Publisher    string    `json:"publisher"`
}

func GetJournalHandle(context *gin.Context) {
	journalIDStr := context.Query("id")
	// 将journalIDStr转换为int64
	journalID, err := strconv.ParseInt(journalIDStr, 10, 64)
	if err != nil {
		// journalID无法被转换为int64, 返回客户端错误
		context.JSON(http.StatusBadRequest, OneJournalResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), "Invalid journal ID")})
		return
	}

	journalService := service.NewJournalService()
	journal, err := journalService.Get(model.Journal{ID: journalID})
	if err != nil {
		// 处理服务层可能返回的错误
		context.JSON(http.StatusInternalServerError, OneJournalResponse{Response: utils.NewCommonResponse(int(enum.OperateFail), err.Error())})
		return
	}

	// 返回信息
	context.JSON(http.StatusOK, OneJournalResponse{
		Journal: Journal{
			Abbreviation: journal.Abbreviation,
			CCFRanking:   journal.CCFRanking,
			Deadline:     journal.Deadline,
			Description:  journal.Description,
			FullName:     journal.FullName,
			ID:           journal.ID,
			ImpactFactor: journal.ImpactFactor,
			ISSN:         journal.ISSN,
			Publisher:    journal.Publisher,
		},
		Response: utils.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
