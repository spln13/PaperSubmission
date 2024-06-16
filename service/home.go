package service

import (
	"PaperSubmission/cache"
	"PaperSubmission/utils"
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
	conferenceNum, _ := NewConferenceService().GetConferenceNum()
	journalNum, _ := NewJournalService().GetJournalNum()
	userNum, _ := NewUserService().GetUserNum()
	_ = cache.NewHomePageCache().SetUserNum(userNum)
	_ = cache.NewHomePageCache().SetJournalNum(journalNum)
	_ = cache.NewHomePageCache().SetConferenceNum(conferenceNum)
}

func UpdateCachedHomeListPeriodically() {
	ticker := time.NewTicker(24 * time.Hour) // 定时一天
	for range ticker.C {
		updateCachedHomeList()
	}
}

func updateCachedHomeList() {
	// 将会议和期刊以及special_issue的前十条存储在cache中
	request := utils.ListQuery{Page: 1, PageSize: 10}
	conferenceList, _ := NewConferenceService().GetList(&request)
	journalList, _ := NewJournalService().GetList(&request)
	specialIssueList, _ := NewSpecialIssueService().GetList(&request)
	_ = cache.NewHomePageCache().CacheConferenceList(conferenceList)
	_ = cache.NewHomePageCache().CacheJournalList(journalList)
	_ = cache.NewHomePageCache().CacheSpecialIssueList(specialIssueList)
}
