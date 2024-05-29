package service

import (
	"PaperSubmission/cache"
	"time"
)

// UpdateHomeInformationPeriodically 一个goroutine负责运行该函数，每一分钟查询数据库将主页参数更新到cache
func UpdateHomeInformationPeriodically() {
	ticker := time.NewTicker(1 * time.Minute) // Update every minute
	for range ticker.C {
		updateHomeInformation()
	}
}

func updateHomeInformation() {
	conferenceNum, _ := NewJournalService().GetJournalNum()
	journalNum, _ := NewJournalService().GetJournalNum()
	userNum, _ := NewUserService().GetUserNum()
	_ = cache.NewHomePageCache().SetUserNum(userNum)
	_ = cache.NewHomePageCache().SetJournalNum(journalNum)
	_ = cache.NewHomePageCache().SetConferenceNum(conferenceNum)
}
