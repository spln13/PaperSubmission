package model

import (
	"PaperSubmission/utils"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type FollowConference struct {
	ID           int64
	UserID       int64     `gorm:"column:user_id"`
	ConferenceID int64     `gorm:"column:conference_id"`
	IsDelete     bool      `gorm:"column:is_delete"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type FollowConferenceModelInterface interface {
	Add(followConference FollowConference) error
	Delete(followConference FollowConference) error
	GetJournalList(userID int64, request utils.ListQuery) ([]int64, error)    // 获取用户关注的所有期刊id,
	GetUserList(conferenceID int64, request utils.ListQuery) ([]int64, error) // 获取关注该期刊的所有用户id
}

type FollowConferenceModel struct{}

var (
	followConferenceModel     *FollowConferenceModel
	followConferenceModelOnce sync.Once
)

func NewFollowConferenceModel() *FollowConferenceModel {
	followConferenceModelOnce.Do(func() {
		followConferenceModel = new(FollowConferenceModel)
	})
	return followConferenceModel
}

func (f FollowConferenceModel) Add(followConference FollowConference) error {
	if err := GetDB().Create(&followConference).Error; err != nil {
		log.Println(err.Error())
		return errors.New("添加记录错误")
	}
	return nil
}

func (f FollowConferenceModel) Delete(followConference FollowConference) error {
	userID := followConference.UserID
	conferenceID := followConference.ConferenceID
	if err := GetDB().Model(&FollowConference{}).Where("user_id=? and journal_id=?", userID, conferenceID).Update("is_delete", true).Error; err != nil {
		log.Println(err)
		return errors.New("删除错误")
	}
	return nil
}

func (f FollowConferenceModel) GetConferenceList(userID int64, request utils.ListQuery) ([]int64, error) {
	var conferenceIDs []int64
	limit, offset := utils.Page(request.PageSize, request.Page) // 分页
	if err := GetDB().Model(&FollowConference{}).Order("id desc").Limit(limit).Offset(offset).Where("user_id=?", userID).Pluck("conference_id", &conferenceIDs).Error; err != nil {
		log.Println(err)
		return []int64{}, nil
	}
	fmt.Println(conferenceIDs)
	return conferenceIDs, nil
}

func (f FollowConferenceModel) GetUserList(conferenceID int64, request utils.ListQuery) ([]int64, error) {
	var userIDs []int64
	limit, offset := utils.Page(request.PageSize, request.Page) // 分页
	if err := GetDB().Model(&FollowConference{}).Order("id desc").Limit(limit).Offset(offset).Where("conference_id=?", conferenceID).Pluck("conference_id", &userIDs).Error; err != nil {
		log.Println(err)
		return []int64{}, nil
	}
	return userIDs, nil
}