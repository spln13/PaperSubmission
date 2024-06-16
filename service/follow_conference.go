package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"sync"
)

type FollowConferenceServiceInterface interface {
	Add(followConference model.FollowConference) error
	Delete(followConference model.FollowConference) error
	Exists(followConference model.FollowConference) (bool, error)
	GetConferenceList(userID int64, request utils.ListQuery) ([]*model.Conference, error) // 获取用户关注的所有期刊id,
	GetUserList(journalID int64, request utils.ListQuery) ([]*model.User, error)          // 获取关注该期刊的所有用户id
}

type FollowConferenceService struct {
}

var (
	followConferenceService     *FollowConferenceService
	followConferenceServiceOnce sync.Once
)

func NewFollowConferenceService() *FollowConferenceService {
	followConferenceServiceOnce.Do(func() {
		followConferenceService = new(FollowConferenceService)
	})
	return followConferenceService
}

func (f FollowConferenceService) Add(followConference model.FollowConference) error {
	return model.NewFollowConferenceModel().Add(followConference)
}

func (f FollowConferenceService) Delete(followConference model.FollowConference) error {
	return model.NewFollowConferenceModel().Delete(followConference)
}

func (f FollowConferenceService) Exists(followConference model.FollowConference) (bool, error) {
	return model.NewFollowConferenceModel().Exist(followConference)
}

func (f FollowConferenceService) GetConferenceList(userID int64, request utils.ListQuery) ([]*model.Conference, error) {
	conferenceIDs, err := model.NewFollowConferenceModel().GetConferenceList(userID, request)
	if err != nil {
		return nil, err
	}
	conferences, err := NewConferenceService().GetSpecifiedList(conferenceIDs) //根据期刊id获取期刊详细信息
	if err != nil {
		return nil, err
	}
	return conferences, nil
}

func (f FollowConferenceService) GetUserList(conferenceID int64, request utils.ListQuery) ([]*model.User, error) {
	userIDs, err := model.NewFollowConferenceModel().GetUserList(conferenceID, request)
	if err != nil {
		return nil, err
	}
	users, err := NewUserService().GetSpecifiedList(userIDs)
	if err != nil {
		return nil, err
	}
	return users, nil
}
