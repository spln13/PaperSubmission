// Package config define parameters needed in this project
package config

var (
	MysqlUsername = "root"
	MysqlPassword = "spln13spln"
	MysqlHost     = "127.0.0.1"
	MysqlPort     = 3306
	MysqlDBName   = "paper_submission"
)

var (
	RedisAddr     = "localhost:6379" // Redis运行地址
	RedisPassword = ""
	RedisDB       = 0 // use default DB
)
