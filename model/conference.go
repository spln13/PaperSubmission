package model

import (
	"PaperSubmission/utils"
	"errors"
	"log"
	"sync"
	"time"
)

type Conference struct {
	ID               int64     `gorm:"primary_key"`
	FullName         string    `gorm:"column:full_name"`
	Link             string    `gorm:"column:link"`
	Abbreviation     string    `gorm:"column:abbreviation"`
	CCFRanking       string    `gorm:"column:ccf_ranking"`
	MeetingVenue     string    `gorm:"column:meeting_venue"`
	Info             string    `gorm:"column:info"`
	Sessions         int64     `gorm:"column:sessions"`
	MaterialDeadline time.Time `gorm:"column:material_deadline"`
	NotificationDate time.Time `gorm:"column:notification_date"`
	MeetingDate      time.Time `gorm:"column:meeting_date"`
	IsDeleted        bool      `gorm:"column:is_deleted"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type ConferenceModelInterface interface {
	Get(conference Conference) (*Conference, error)
	GetList(request *utils.ListQuery) ([]*Conference, error)
	GetConferenceNum() (int64, error)
	GetSpecifiedList(conferenceIDs []int64) ([]*Conference, error)
}

type ConferenceModel struct {
}

var (
	conferenceModel *ConferenceModel
	conferenceOnce  sync.Once
)

func NewConferenceModel() *ConferenceModel {
	conferenceOnce.Do(func() {
		conferenceModel = new(ConferenceModel)
	})
	return conferenceModel
}

func (c ConferenceModel) Get(conference Conference) (*Conference, error) {
	if err := GetDB().Model(&conference).
		Select("id", "full_name", "link", "abbreviation", "ccf_ranking", "meeting_venue", "info", "sessions", "material_deadline", "notification_date", "meeting_date").First(&conference).Error; err != nil {
		log.Println(err)
		return nil, errors.New("查询会议信息错误")
	}
	return &conference, nil
}

func (c ConferenceModel) GetList(request *utils.ListQuery) ([]*Conference, error) {
	var conferences []*Conference
	today := time.Now()
	limit, offset := request.PageSize, request.PageSize // 分页
	if err := GetDB().Where("material_deadline > ?", today).Order("material_deadline asc").Limit(limit).Offset(offset).Select("id", "full_name", "link", "abbreviation", "ccf_ranking", "meeting_venue", "info", "sessions", "material_deadline", "notification_date", "meeting_date").Find(&conferences).Error; err != nil {
		log.Println(err)
		return nil, errors.New("查询会议信息错误")
	}
	return conferences, nil
}

func (c ConferenceModel) GetSpecifiedList(conferenceIDs []int64) ([]*Conference, error) {
	var conferences []*Conference
	if err := GetDB().Where("id in (?)", conferenceIDs).Select("id", "full_name", "link", "abbreviation", "ccf_ranking", "meeting_venue", "info", "sessions", "material_deadline", "notification_date", "meeting_date").Find(&conferences).Error; err != nil {
		return nil, errors.New("获取会议信息错误")
	}
	return conferences, nil
}

func (c ConferenceModel) GetConferenceNum() (int64, error) {
	var count int64
	if err := GetDB().Model(&Conference{}).Count(&count).Error; err != nil {
		log.Println(err)
		return -1, errors.New("查询会议数量错误")
	}
	return count, nil
}
