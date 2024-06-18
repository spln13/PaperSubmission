package controller

import (
	"PaperSubmission/enum"
	"PaperSubmission/model"
	"PaperSubmission/response"
	"PaperSubmission/service"
	"PaperSubmission/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type OneJournalResponse struct {
	Journal
	response.Response
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
	Link         string    `json:"link"`
	Publisher    string    `json:"publisher"`
}

type JournalListResponse struct {
	List []Journal `json:"list"`
	response.Response
}

func GetJournalHandle(context *gin.Context) {
	journalIDStr := context.Query("id")
	// 将journalIDStr转换为int64
	journalID, err := strconv.ParseInt(journalIDStr, 10, 64)
	if err != nil {
		// journalID无法被转换为int64, 返回客户端错误
		context.JSON(http.StatusBadRequest, OneJournalResponse{Response: response.NewCommonResponse(int(enum.OperateFail), "Invalid journal ID")})
		return
	}

	journalService := service.NewJournalService()
	journal, err := journalService.Get(model.Journal{ID: journalID})
	if err != nil {
		// 处理服务层可能返回的错误
		context.JSON(http.StatusInternalServerError, OneJournalResponse{Response: response.NewCommonResponse(int(enum.OperateFail), err.Error())})
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
			Link:         journal.Link,
			ID:           journal.ID,
			ImpactFactor: journal.ImpactFactor,
			ISSN:         journal.ISSN,
			Publisher:    journal.Publisher,
		},
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func JournalListHandler(context *gin.Context) {
	pageStr := context.Query("page")
	pageSizeStr := context.Query("page_size")
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	fmt.Println(pageStr, pageSizeStr)
	if err1 != nil || err2 != nil { // 请求参数无法被解析
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	request := utils.ListQuery{Page: page, PageSize: pageSize}
	journalModelList, err := service.NewJournalService().GetList(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var journalList []Journal
	for _, journalModel := range journalModelList {
		journal := Journal{
			Abbreviation: journalModel.Abbreviation,
			CCFRanking:   journalModel.CCFRanking,
			Deadline:     journalModel.Deadline,
			Description:  journalModel.Description,
			FullName:     journalModel.FullName,
			ID:           journalModel.ID,
			Link:         journalModel.Link,
			ImpactFactor: journalModel.ImpactFactor,
			ISSN:         journalModel.ISSN,
			Publisher:    journalModel.Publisher,
		}
		journalList = append(journalList, journal)
	}
	context.JSON(http.StatusOK, JournalListResponse{
		List:     journalList,
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func QueryJournalHandler(context *gin.Context) {
	key := context.Query("key")
	journalModelList, err := service.NewJournalService().Query(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	var journalList []Journal
	for _, journalModel := range journalModelList {
		journal := Journal{
			Abbreviation: journalModel.Abbreviation,
			CCFRanking:   journalModel.CCFRanking,
			Deadline:     journalModel.Deadline,
			Description:  journalModel.Description,
			FullName:     journalModel.FullName,
			ID:           journalModel.ID,
			Link:         journalModel.Link,
			ImpactFactor: journalModel.ImpactFactor,
			ISSN:         journalModel.ISSN,
			Publisher:    journalModel.Publisher,
		}
		journalList = append(journalList, journal)
	}
	context.JSON(http.StatusOK, JournalListResponse{
		List:     journalList,
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
