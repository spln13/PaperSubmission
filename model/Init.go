package model

import (
	"PaperSubmission/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Init() {
	// configure parameters for mysql
	username := config.MysqlUsername
	password := config.MysqlPassword
	host := config.MysqlHost
	port := config.MysqlPort
	DBName := config.MysqlDBName
	//dsn := "root:spln13spln@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, DBName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // print sql sentences
	})
	if err != nil { // fail to connect with mysql, need to use panic
		panic(err)
	}
	// create tables in database

	//err = sysDB.AutoMigrate(&AdminAccount{}, &TeacherAccount{}, &StudentAccount{}, &ExerciseAssociation{},
	//	&ExerciseTable{}, &ExerciseContent{}, &SubmitHistory{}, &UserProblemStatus{}, &Contest{}, &Class{},
	//	&ContestClassAssociation{}, &ContestExerciseAssociation{}, &ScoreRecord{}, &ContestSubmission{}, &ContestExerciseStatus{})
	err = db.AutoMigrate(&User{}) // 建表
	if err != nil {
		panic(err)
	}
	db, err := db.DB()
	if err != nil {
		log.Fatalln("db connected error", err)
	}
	//db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
}
