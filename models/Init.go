package models

import (
	"PaperSubmission/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var sysDB *gorm.DB

func InitSysDB() {
	// configure parameters for mysql
	username := config.MysqlUsername
	password := config.MysqlPassword
	host := config.MysqlHost
	port := config.MysqlPort
	SysDBName := config.MysqlDBName
	//dsn := "root:spln13spln@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, SysDBName)
	var err error
	sysDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // print sql sentences
	})
	if err != nil { // fail to connect with mysql, need to use panic
		panic(err)
	}
	// create tables in database

	//err = sysDB.AutoMigrate(&AdminAccount{}, &TeacherAccount{}, &StudentAccount{}, &ExerciseAssociation{},
	//	&ExerciseTable{}, &ExerciseContent{}, &SubmitHistory{}, &UserProblemStatus{}, &Contest{}, &Class{},
	//	&ContestClassAssociation{}, &ContestExerciseAssociation{}, &ScoreRecord{}, &ContestSubmission{}, &ContestExerciseStatus{})
	if err != nil {
		panic(err)
	}

	db, err := sysDB.DB()
	if err != nil {
		log.Fatalln("db connected error", err)
	}
	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
}
