package main

import (
	"PaperSubmission/cache"
	"PaperSubmission/model"
	"PaperSubmission/service"
)

func main() {
	server := InitServer()
	model.Init()
	cache.InitRedis()
	go service.UpdateHomeInformationPeriodically()
	go service.UpdateCachedHomeListPeriodically()
	//go service.UpdateConferenceDatabasePeriodically()
	//go service.UpdateJournalDatabasePeriodically()
	if err := server.Run(":8080"); err != nil {
		// 运行在8080端口
		panic(err)
	}

}
