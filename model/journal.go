package model

import (
	"PaperSubmission/utils"
	"errors"
	"log"
	"sync"
	"time"
)

type Journal struct {
	ID           int64     `gorm:"primary_key"`
	FullName     string    `gorm:"column:full_name"`
	Link         string    `gorm:"column:link"`
	ImpactFactor float64   `gorm:"column:impact_factor"`
	Abbreviation string    `gorm:"column:abbreviation"`
	Publisher    string    `gorm:"column:publisher"`
	ISSN         string    `gorm:"column:issn"`
	Description  string    `gorm:"column:description"`
	CCFRanking   string    `gorm:"column:ccf_ranking"`
	Deadline     time.Time `gorm:"column:deadline"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type JournalModelInterface interface {
	GetList(request *utils.ListQuery) ([]*Journal, error)
	Get(journal Journal) (*Journal, error)
	Query(key string) ([]*Journal, error)
	GetSpecifiedList(journalIDs []int64) ([]*Journal, error)
	GetJournalNum() (int64, error)
}

type JournalModel struct {
}

var (
	journalModel     *JournalModel
	journalModelOnce sync.Once
)

func NewJournalModel() *JournalModel {
	journalModelOnce.Do(func() {
		journalModel = new(JournalModel)
	})
	return journalModel
}

func (j JournalModel) GetList(request *utils.ListQuery) ([]*Journal, error) {
	var journals []*Journal
	limit, offset := utils.Page(request.PageSize, request.Page)
	if err := GetDB().Order("impact_factor desc").Limit(limit).Offset(offset).Select("id", "full_name", "link", "impact_factor", "abbreviation", "publisher", "issn", "description", "ccf_ranking", "deadline").Find(&journals).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取期刊信息错误")
	}
	return journals, nil
}

func (j JournalModel) Get(journal Journal) (*Journal, error) {
	if err := GetDB().Model(&journal).Select("id", "full_name", "link", "impact_factor", "abbreviation", "publisher", "issn", "description", "ccf_ranking", "deadline").First(&journal).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取期刊信息错误")
	}
	return &journal, nil
}

func (j JournalModel) Query(key string) ([]*Journal, error) {
	var journals []*Journal
	if err := GetDB().Where("full_name LIKE ? OR abbreviation LIKE ?", "%"+key+"%", "%"+key+"%").Find(&journals).Error; err != nil {
		log.Println(err)
		return nil, errors.New("搜索会议信息错误")
	}
	return journals, nil
}

func (j JournalModel) GetSpecifiedList(journalIDs []int64) ([]*Journal, error) {
	var journals []*Journal
	if err := GetDB().Where("id in (?)", journalIDs).Select("id", "full_name", "link", "impact_factor", "abbreviation", "publisher", "issn", "description", "ccf_ranking", "deadline").Find(&journals).Error; err != nil {
		return nil, errors.New("获取期刊信息错误")
	}
	return journals, nil
}

func (j JournalModel) GetJournalNum() (int64, error) {
	var count int64
	if err := GetDB().Model(&Journal{}).Count(&count).Error; err != nil {
		log.Println(err.Error())
		return -1, errors.New("查询期刊数量错误")
	}
	return count, nil
}
