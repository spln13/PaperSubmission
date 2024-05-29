package cache

import (
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
