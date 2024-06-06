package main

import (
	"PaperSubmission/model"
)

func main() {
	server := InitServer()
	model.Init()
	//	cache.InitRedis()
	//go service.UpdateHomeInformationPeriodically()
	if err := server.Run(":8080"); err != nil {
		// 运行在8080端口
		panic(err)
	}

}
