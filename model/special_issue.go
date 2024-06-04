package model

import (
	"PaperSubmission/utils"
	"errors"
	"log"
	"sync"
	"time"
)

type SpecialIssue struct {
	ID           int64     `gorm:"column:id"`
	JournalID    int64     `gorm:"column:journal_id"`
	IssueContent string    `gorm:"column:issue_content"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type SpecialIssueModelInterface interface {
	Get(specialIssue *SpecialIssue) (*SpecialIssue, error)
	GetList(request *utils.ListQuery) ([]*SpecialIssue, error)
}

type SpecialIssueModel struct {
}

var (
	specialIssueModel     *SpecialIssueModel
	specialIssueModelOnce sync.Once
)

func NewSpecialIssueModel() *SpecialIssueModel {
	specialIssueModelOnce.Do(func() {
		specialIssueModel = new(SpecialIssueModel)
	})
	return specialIssueModel
}

func (s SpecialIssueModel) Get(specialIssue *SpecialIssue) (*SpecialIssue, error) {
	if err := GetDB().Model(&specialIssue).Select("id", "journal_id", "issue_content").Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取special issue信息错误")
	}
	return specialIssue, nil
}

func (s SpecialIssueModel) GetList(request *utils.ListQuery) ([]*SpecialIssue, error) {
	var specialIssues []*SpecialIssue
	limit, offset := utils.Page(request.PageSize, request.Page)
	if err := GetDB().Order("id desc").Limit(limit).Offset(offset).Select("id", "journal_id", "issue_content").Find(&specialIssues).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取special issue信息错误")
	}
	return specialIssues, nil
}
