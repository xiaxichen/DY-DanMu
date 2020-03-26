package mysql

import (
	"DY-DanMu/dbMysql/config"
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Log "github.com/sirupsen/logrus"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", config.MysqlUser+":"+config.MysqlPWD+"@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		Log.Errorf("Failed to connect to mysql,err:%s", err)
		os.Exit(1)
	}
	err = __createDB(db)
	if err != nil {
		panic(err)
	}
	err = __createTable(db)
	if err != nil {
		panic(err)
	}
}

//__createDB:创建DB
func __createDB(conn *sql.DB) error {
	prepare, err := conn.Prepare(fmt.Sprintf(`create database if not exists %s;`, config.MysqlDBName))
	if err != nil {
		return err
	}
	defer prepare.Close()
	exec, err := prepare.Exec()
	if err != nil {
		return err
	}
	if affected, err := exec.RowsAffected(); err == nil && affected > 0 {
		return nil
	} else {
		return err
	}
}

//__createTable:创建表
func __createTable(conn *sql.DB) error {
	prepare, err := conn.Prepare(fmt.Sprintf(`
CREATE TABLE  IF NOT EXISTS dybarrage.%s  (
  cid varchar(32) NOT NULL,
  level int(10) NOT NULL,
  bl int(32) NULL,
  bnn varchar(32) NULL,
  brid varchar(32) NULL,
  col varchar(32) NULL,
  cst bigint(13) NOT NULL,
  dms varchar(32) NULL,
  ifs varchar(32) NULL,
  esIndex varchar(32) NOT NULL,
  hc varchar(32) NULL,
  urlev int(10) NULL,
  type varchar(32) NULL,
  sahf varchar(32) NULL,
  lk varchar(32) NULL,
  fl varchar(32) NULL,
  el varchar(32) NULL,
  ct varchar(32) NULL,
  txt varchar(255) NULL,
  nn varchar(32) NOT Null,
  uid int(32) NOT NULL,
  PRIMARY KEY (cid),
  KEY cst (cst) USING BTREE COMMENT '时间索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;`, config.MysqlTableName))
	if err != nil {
		return err
	}
	defer prepare.Close()
	exec, err := prepare.Exec()
	if _, err := exec.RowsAffected(); err == nil {
		return nil
	} else {
		return err
	}
}

//DBConn:返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
func checkErr(err error) {
	if err != nil {
		Log.Fatal(err)
		panic(err)
	}
}

// ParseRows:解析数据到行中
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}
