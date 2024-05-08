package model

import (
	"sync"
	"time"
)

type Journal struct {
	ID           int64     `gorm:"primary_key"`
	Name         string    `gorm:"column:name"`
	SpecialIssue string    `gorm:"column:special_issue"`
	Link         string    `gorm:"column:link"`
	ImpactFactor float64   `gorm:"column:impact_factor"`
	Publisher    string    `gorm:"column:publisher"`
	ISSN         string    `gorm:"column:issn"`
	Description  string    `gorm:"column:description"`
	CCFRanking   string    `gorm:"column:ccf_ranking"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	Deadline     time.Time `gorm:"column:deadline"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type JournalModelInterface interface {
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
