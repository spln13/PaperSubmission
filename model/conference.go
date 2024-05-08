package model

import "time"

type Conference struct {
	ID           int64     `gorm:"primary_key"`
	FullName     string    `gorm:"full_name"`
	Abbreviation string    `gorm:"abbreviation"`
	Location     string    `gorm:"location"`
	CCFRanking   string    `gorm:"column:ccf_ranking"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	Deadline     time.Time `gorm:"column:deadline"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
