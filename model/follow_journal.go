package model

import (
	"PaperSubmission/utils"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type FollowJournal struct {
	ID        int64
	UserID    int64     `gorm:"column:user_id"`
	JournalID int64     `gorm:"column:journal_id"`
	IsDelete  bool      `gorm:"column:is_delete"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type FollowJournalModelInterface interface {
	Add(followJournal FollowJournal) error
	Delete(followJournal FollowJournal) error
	GetJournalList(userID int64, request utils.ListQuery) ([]int64, error) // 获取用户关注的所有期刊id,
	GetUserList(journalID int64, request utils.ListQuery) ([]int64, error) // 获取关注该期刊的所有用户id
}

type FollowJournalModel struct{}

var (
	followJournalModel     *FollowJournalModel
	followJournalModelOnce sync.Once
)

func NewFollowJournalModel() *FollowJournalModel {
	followJournalModelOnce.Do(func() {
		followJournalModel = new(FollowJournalModel)
	})
	return followJournalModel
}

func (f FollowJournalModel) Add(followJournal FollowJournal) error {
	if err := GetDB().Create(&followJournal).Error; err != nil {
		log.Println(err.Error())
		return errors.New("添加记录错误")
	}
	return nil
}

func (f FollowJournalModel) Delete(followJournal FollowJournal) error {
	userID := followJournal.UserID
	journalID := followJournal.JournalID
	if err := GetDB().Model(&FollowJournal{}).Where("user_id=? and journal_id=?", userID, journalID).Update("is_delete", true).Error; err != nil {
		log.Println(err)
		return errors.New("删除错误")
	}
	return nil
}

func (f FollowJournalModel) GetJournalList(userID int64, request utils.ListQuery) ([]int64, error) {
	var journalIDs []int64
	limit, offset := utils.Page(request.PageSize, request.Page) // 分页
	if err := GetDB().Model(&FollowJournal{}).Order("id desc").Limit(limit).Offset(offset).Where("user_id=?", userID).Pluck("journal_id", &journalIDs).Error; err != nil {
		log.Println(err)
		return []int64{}, nil
	}
	fmt.Println(journalIDs)
	return journalIDs, nil
}

func (f FollowJournalModel) GetUserList(journalID int64, request utils.ListQuery) ([]int64, error) {
	var userIDs []int64
	limit, offset := utils.Page(request.PageSize, request.Page) // 分页
	if err := GetDB().Model(&FollowJournal{}).Order("id desc").Limit(limit).Offset(offset).Where("journal_id=?", journalID).Pluck("journal_id", &userIDs).Error; err != nil {
		log.Println(err)
		return []int64{}, nil
	}
	return userIDs, nil
}
