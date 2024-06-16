package service

import (
	"PaperSubmission/model"
	"PaperSubmission/response"
	"PaperSubmission/utils"
	"sync"
)

type SpecialIssueServiceInterface interface {
	Get(specialIssue response.SpecialIssue) (response.SpecialIssue, error)
	GetList(request *utils.ListQuery) ([]response.SpecialIssue, error)
	GetModelList(request *utils.ListQuery) ([]model.SpecialIssue, error)
}

var (
	specialIssueService     *SpecialIssueService
	specialIssueServiceOnce sync.Once
)

type SpecialIssueService struct{}

func NewSpecialIssueService() *SpecialIssueService {
	specialIssueServiceOnce.Do(func() {
		specialIssueService = new(SpecialIssueService)
	})
	return specialIssueService
}

func (s SpecialIssueService) Get(specialIssue response.SpecialIssue) (response.SpecialIssue, error) {
	specialIssueModel := &model.SpecialIssue{
		ID: specialIssue.ID,
	}
	getSpecialIssueModel, err := model.NewSpecialIssueModel().Get(specialIssueModel)
	if err != nil {
		return response.SpecialIssue{}, err
	}
	specialIssue.IssueContent = getSpecialIssueModel.IssueContent
	journalID := getSpecialIssueModel.JournalID
	journalModel := model.Journal{ID: journalID}
	getJournalModel, err := model.NewJournalModel().Get(journalModel)
	if err != nil {
		return response.SpecialIssue{}, err
	}
	specialIssue.JournalID = getJournalModel.ID
	specialIssue.FullName = getJournalModel.FullName
	specialIssue.Publisher = getJournalModel.Publisher
	specialIssue.Description = getJournalModel.Description
	specialIssue.Issn = getJournalModel.ISSN
	specialIssue.ImpactFactor = getJournalModel.ImpactFactor
	specialIssue.Abbreviation = getJournalModel.Abbreviation
	return specialIssue, nil
}

func (s SpecialIssueService) GetList(request *utils.ListQuery) ([]response.SpecialIssue, error) {
	specialIssueModelList, err := model.NewSpecialIssueModel().GetList(request) // id, journalID, content
	if err != nil {
		return nil, err
	}
	var specialIssueList []response.SpecialIssue
	for _, specialIssueModel := range specialIssueModelList {
		journalID := specialIssueModel.JournalID
		journalModel := model.Journal{ID: journalID}
		getJournalModel, _ := model.NewJournalModel().Get(journalModel)
		var specialIssue response.SpecialIssue
		specialIssue.ID = specialIssueModel.ID
		specialIssue.IssueContent = specialIssueModel.IssueContent
		specialIssue.FullName = getJournalModel.FullName
		specialIssue.Publisher = getJournalModel.Publisher
		specialIssue.Description = getJournalModel.Description
		specialIssue.Issn = getJournalModel.ISSN
		specialIssue.ImpactFactor = getJournalModel.ImpactFactor
		specialIssue.Abbreviation = getJournalModel.Abbreviation
		specialIssue.CcfRanking = getJournalModel.CCFRanking
		specialIssueList = append(specialIssueList, specialIssue)
	}
	return specialIssueList, nil
}

func (s SpecialIssueService) GetModelList(request *utils.ListQuery) ([]*model.SpecialIssue, error) {
	return model.NewSpecialIssueModel().GetList(request)
}
