package config

import "os"

const (
	MysqlUser      = "root"
	MysqlDBName    = "dybarrage"
	MysqlTableName = "barrage"
	RedisAddr      = "127.0.0.1:6379"
	RedisDB        = 0
)

var (
	MysqlPWD string
	RedisPWD string
)

func init() {
	MysqlPWD = os.Getenv("MYSQLPWD")
	RedisPWD = os.Getenv("REDISPWD")
}
