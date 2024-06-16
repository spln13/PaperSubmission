package cache

import (
	"PaperSubmission/model"
	"PaperSubmission/response"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"sync"
)

type HomePageInterface interface {
	IncreasePageView() error
	GetConferenceNum() (int64, error)
	GetJournalNum() (int64, error)
	GetUserNum() (int64, error)
	GetPageView() (int64, error)
	SetConferenceNum(meetingNum int64) error
	SetJournalNum(journalNum int64) error
	SetUserNum(userNum int64) error
	CacheConferenceList(conferenceLust []*model.Conference) error
	CacheJournalList(journalList []*model.Journal) error
	CacheSpecialIssueList(specialIssue []response.SpecialIssue) error
	GetConferenceList() ([]model.Conference, error)
	GetJournalList() ([]model.Journal, error)
	GetSpecialIssueList() ([]response.SpecialIssue, error)
}

type HomePageCache struct {
}

var (
	homePageCache     *HomePageCache
	homePageCacheOnce sync.Once
)

func NewHomePageCache() *HomePageCache {
	homePageCacheOnce.Do(func() {
		homePageCache = new(HomePageCache)
	})
	return homePageCache
}

func (h HomePageCache) IncreasePageView() error {
	key := "page_view"
	if err := rdb.Incr(ctx, key).Err(); err != nil { // 页面浏览自增
		log.Printf("Error increasing page view: %v\n", err)
		return errors.New("error increasing page view")
	}
	return nil
}

func (h HomePageCache) GetConferenceNum() (int64, error) {
	key := "meeting_num"
	numStr, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting meeting num: %v\n", err)
		return -1, errors.New("error getting meeting num")
	}
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num, nil
}

func (h HomePageCache) GetJournalNum() (int64, error) {
	key := "journal_num"
	numStr, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting journal num: %v\n", err)
		return -1, errors.New("error getting journal num")

	}
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num, nil
}

func (h HomePageCache) GetUserNum() (int64, error) {
	key := "user_num"
	numStr, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting user num: %v\n", err)
		return -1, errors.New("error getting user num")
	}
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num, nil
}

func (h HomePageCache) SetConferenceNum(meetingNum int64) error {
	key := "meeting_num"
	meetingNumStr := strconv.FormatInt(meetingNum, 10)
	if err := rdb.Set(ctx, key, meetingNumStr, 0).Err(); err != nil {
		log.Printf("Error setting meeting num: %v\n", err)
		return errors.New("error setting meeting num")
	}
	return nil
}

func (h HomePageCache) SetUserNum(userNum int64) error {
	key := "user_num"
	userNumStr := strconv.FormatInt(userNum, 10)
	if err := rdb.Set(ctx, key, userNumStr, 0).Err(); err != nil {
		log.Printf("Error setting user num: %v\n", err)
		return errors.New("error setting user num")
	}
	return nil
}

func (h HomePageCache) SetJournalNum(journalNum int64) error {
	key := "journal_num"
	journalNumStr := strconv.FormatInt(journalNum, 10)
	if err := rdb.Set(ctx, key, journalNumStr, 0).Err(); err != nil {
		log.Printf("Error setting journal num: %v\n", err)
		return errors.New("error setting journal num")
	}
	return nil
}

func (h HomePageCache) GetPageView() (int64, error) {
	key := "page_view"
	numStr, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting page view: %v\n", err)
		return -1, errors.New("error getting page view")
	}
	num, _ := strconv.ParseInt(numStr, 10, 64)
	return num, nil
}

func (h HomePageCache) CacheConferenceList(conferenceLust []*model.Conference) error {
	jsonData, err := json.Marshal(conferenceLust)
	if err != nil {
		log.Printf("Error occurred during marshalling. %v\n", err)
		return errors.New("error occurred during marshalling")
	}
	err = rdb.Set(ctx, "conference_list", jsonData, 0).Err()
	if err != nil {
		log.Printf("Error occurred during saving to Redis. %v", err)
		return errors.New("error occurred during saving to Redis")
	}
	return nil
}

func (h HomePageCache) CacheJournalList(journalList []*model.Journal) error {
	jsonData, err := json.Marshal(journalList)
	if err != nil {
		log.Printf("Error occurred during marshalling. %v\n", err)
		return errors.New("error occurred during marshalling")
	}
	err = rdb.Set(ctx, "journal_list", jsonData, 0).Err()
	if err != nil {
		log.Printf("Error occurred during saving to Redis. %v", err)
		return errors.New("error occurred during saving to Redis")
	}
	return nil
}

func (h HomePageCache) CacheSpecialIssueList(specialIssue []response.SpecialIssue) error {
	jsonData, err := json.Marshal(specialIssue)
	if err != nil {
		log.Printf("Error occurred during marshalling. %v\n", err)
		return errors.New("error occurred during marshalling")
	}
	err = rdb.Set(ctx, "special_issue_list", jsonData, 0).Err()
	if err != nil {
		log.Printf("Error occurred during saving to Redis. %v", err)
		return errors.New("error occurred during saving to Redis")
	}
	return nil
}

func (h HomePageCache) GetConferenceList() ([]model.Conference, error) {
	val, err := rdb.Get(ctx, "conference_list").Bytes()
	if err != nil {
		log.Printf("error occurred during retrieving from Redis. %v\n", err)
		return nil, errors.New("error occurred during retrieving from Redis")
	}
	var conferenceList []model.Conference
	err = json.Unmarshal(val, &conferenceList)
	if err != nil {
		log.Printf("error occurred during unmarshalling from Redis. %v\n", err)
		return nil, errors.New("error occurred during unmarshalling from Redis")
	}
	return conferenceList, nil
}

func (h HomePageCache) GetJournalList() ([]model.Journal, error) {
	val, err := rdb.Get(ctx, "journal_list").Bytes()
	if err != nil {
		log.Printf("error occurred during retrieving from Redis. %v\n", err)
		return nil, errors.New("error occurred during retrieving from Redis")
	}
	var journalList []model.Journal
	err = json.Unmarshal(val, &journalList)
	if err != nil {
		log.Printf("error occurred during unmarshalling from Redis. %v\n", err)
		return nil, errors.New("error occurred during unmarshalling from Redis")
	}
	return journalList, nil
}

func (h HomePageCache) GetSpecialIssueList() ([]response.SpecialIssue, error) {
	val, err := rdb.Get(ctx, "special_issue_list").Bytes()
	if err != nil {
		log.Printf("error occurred during retrieving from Redis. %v\n", err)
		return nil, errors.New("error occurred during retrieving from Redis")
	}
	var specialIssueList []response.SpecialIssue
	err = json.Unmarshal(val, &specialIssueList)
	if err != nil {
		log.Printf("error occurred during unmarshalling from Redis. %v\n", err)
		return nil, errors.New("error occurred during unmarshalling from Redis")
	}
	return specialIssueList, nil
}
