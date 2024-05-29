package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"sync"
)

type ConferenceServiceInterface interface {
	Get(conference model.Conference) (*model.Conference, error)
	GetList(request *utils.ListQuery) ([]*model.Conference, error)
	GetConferenceNum() (int64, error)
}

type ConferenceService struct{}

var (
	conferenceServiceOnce sync.Once
	conferenceService     *ConferenceService
)

func NewConferenceService() *ConferenceService {
	conferenceServiceOnce.Do(func() {
		conferenceService = new(ConferenceService)
	})
	return conferenceService
}

func (c ConferenceService) Get(conference model.Conference) (*model.Conference, error) {
	return model.NewConferenceModel().Get(conference)
}

func (c ConferenceService) GetList(request *utils.ListQuery) ([]*model.Conference, error) {
	return model.NewConferenceModel().GetList(request)
}

func (c ConferenceService) GetConferenceNum() (int64, error) {
	return model.NewConferenceModel().GetConferenceNum()
}
