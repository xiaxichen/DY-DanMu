package config

import "os"

const (
	MysqlUser      = "root"
	MysqlDBName    = "dybarrage"
	MysqlTableName = "barrage"
)

var (
	MysqlPWD string
)

func init() {
	MysqlPWD = os.Getenv("MYSQLPWD")
}
