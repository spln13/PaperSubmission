package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"sync"
)

type JournalServiceInterface interface {
	GetList(request *utils.ListQuery) ([]*model.Journal, error)
	Get(journal model.Journal) (*model.Journal, error)
	Query(key string) ([]*model.Journal, error)
	GetSpecifiedList(journalIDs []int64) ([]*model.Journal, error)
	GetJournalNum() (int64, error)
}

type JournalService struct{}

var (
	journalService     *JournalService
	journalServiceOnce sync.Once
)

func NewJournalService() *JournalService {
	journalServiceOnce.Do(func() {
		journalService = new(JournalService)
	})
	return journalService
}

func (j JournalService) GetList(request *utils.ListQuery) ([]*model.Journal, error) {
	return model.NewJournalModel().GetList(request)
}

func (j JournalService) Get(journal model.Journal) (*model.Journal, error) {
	return model.NewJournalModel().Get(journal)
}

func (j JournalService) Query(key string) ([]*model.Journal, error) {
	return model.NewJournalModel().Query(key)
}

func (j JournalService) GetSpecifiedList(journalIDs []int64) ([]*model.Journal, error) {
	return model.NewJournalModel().GetSpecifiedList(journalIDs)
}

func (j JournalService) GetJournalNum() (int64, error) {
	return model.NewJournalModel().GetJournalNum()
}
