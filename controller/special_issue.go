package controller

import (
	"PaperSubmission/enum"
	"PaperSubmission/response"
	"PaperSubmission/service"
	"PaperSubmission/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OneSpecialIssueResponse struct {
	response.SpecialIssue
	response.Response
}

type SpecialIssueListResponse struct {
	List []response.SpecialIssue `json:"list"`
	response.Response
}

func GetSpecialIssueHandle(context *gin.Context) {
	specialIssueIDStr := context.Query("id")
	specialIssueID, err := strconv.ParseInt(specialIssueIDStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	specialIssue := response.SpecialIssue{ID: specialIssueID}
	specialIssue, err = service.NewSpecialIssueService().Get(specialIssue)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	context.JSON(http.StatusOK, OneSpecialIssueResponse{
		SpecialIssue: specialIssue,
		Response:     response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}

func SpecialIssueListHandler(context *gin.Context) {
	pageStr := context.Query("page")
	pageSizeStr := context.Query("page_size")
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil { // 请求参数无法被解析
		context.JSON(http.StatusBadRequest, response.NewCommonResponse(int(enum.OperateFail), enum.OperateFail.String()))
		return
	}
	request := utils.ListQuery{Page: page, PageSize: pageSize}
	list, err := service.NewSpecialIssueService().GetList(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.NewCommonResponse(int(enum.OperateFail), err.Error()))
		return
	}
	context.JSON(http.StatusOK, SpecialIssueListResponse{
		List:     list,
		Response: response.NewCommonResponse(int(enum.OperateOK), enum.OperateOK.String()),
	})
}
