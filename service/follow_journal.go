package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"sync"
)

type FollowJournalServiceInterface interface {
	Add(followJournal model.FollowJournal) error
	Delete(followJournal model.FollowJournal) error
	GetJournalList(userID int64, request utils.ListQuery) ([]*model.Journal, error) // 获取用户关注的所有期刊id,
	GetUserList(journalID int64, request utils.ListQuery) ([]*model.User, error)    // 获取关注该期刊的所有用户id
}

type FollowJournalService struct {
}

var (
	followJournalService     *FollowJournalService
	followJournalServiceOnce sync.Once
)

func NewFollowJournalService() *FollowJournalService {
	followJournalServiceOnce.Do(func() {
		followJournalService = new(FollowJournalService)
	})
	return followJournalService
}

func (f FollowJournalService) Add(followJournal model.FollowJournal) error {
	return model.NewFollowJournalModel().Add(followJournal)
}

func (f FollowJournalService) Delete(followJournal model.FollowJournal) error {
	return model.NewFollowJournalModel().Delete(followJournal)
}

func (f FollowJournalService) GetJournalList(userID int64, request utils.ListQuery) ([]*model.Journal, error) {
	journalIDs, err := model.NewFollowJournalModel().GetJournalList(userID, request)
	if err != nil {
		return nil, err
	}
	journals, err := NewJournalService().GetSpecifiedList(journalIDs) //根据期刊id获取期刊详细信息
	if err != nil {
		return nil, err
	}
	return journals, nil
}

func (f FollowJournalService) GetUserList(journalID int64, request utils.ListQuery) ([]*model.User, error) {
	userIDs, err := model.NewFollowJournalModel().GetUserList(journalID, request)
	if err != nil {
		return nil, err
	}
	users, err := NewUserService().GetSpecifiedList(userIDs)
	if err != nil {
		return nil, err
	}
	return users, nil
}
