package service

import (
	"PaperSubmission/model"
	"PaperSubmission/utils"
	"sync"
)

type JournalServiceInterface interface {
	GetList(request *utils.ListQuery) ([]*model.Journal, error)
	Get(journal model.Journal) (*model.Journal, error)
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

func (j JournalService) GetJournalNum() (int64, error) {
	return model.NewJournalModel().GetJournalNum()
}
