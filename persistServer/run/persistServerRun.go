package main

import (
	"DY-DanMu/DMconfig/config"
	_type2 "DY-DanMu/dbConn/_type"
	mysqlConfig "DY-DanMu/dbConn/config"
	"DY-DanMu/dbConn/mysql"
	"DY-DanMu/dbConn/redisConn"
	"DY-DanMu/persistServer/rpcsupport"
	"DY-DanMu/persistServer/server"
	"fmt"
	"github.com/olivere/elastic/v7"
	Log "github.com/sirupsen/logrus"
	"net/smtp"
	"strings"
	"time"
)

func main() {
	Log.Fatal(serverRpc(config.SpiderConfig.ItemSaverPort, config.SpiderConfig.ElasticIndex))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time.Minute)
	redisConnInstance := redisConn.DBConn()
	service := &server.ItemSaverService{
		Client:    client,
		Conn:      mysql.DBConn(),
		RedisConn: redisConnInstance,
		Index:     index,
		Count:     0,
		SumCount:  0,
	}
	service1 := &server.SelectMiddlerWare{
		RedisClient: redisConnInstance,
		Client:      client,
		Conn:        mysql.DBConn(),
		Index:       config.DYWebConfig.ElasticIndex,
		Count:       0,
	}
	prepare, err := service.Conn.Prepare(fmt.Sprintf("select count(*) from %s",mysqlConfig.MysqlDBName+"."+mysqlConfig.MysqlTableName))
	if err != nil {
		return err
	}
	defer prepare.Close()
	query, err := prepare.Query()
	if err != nil {
		Log.Error(err)
		return nil
	}
	count := _type2.BarrageCount{}
	for query.Next() {
		err := query.Scan(&count.Count)
		if err != nil {
			return err
		}
		set := service1.RedisClient.Set("BarrageCount", count.Count, 0)
		err = set.Err()
		Log.Info(set)
		if err != nil {
			return err
		}
	}
	serviceEmail := &server.EmialSendSever{
		Auth: smtp.PlainAuth("Send", config.DYWebConfig.EmailUser, config.DYWebConfig.EmailPwd, strings.Split(config.DYWebConfig.EmailHost, ":")[0]),
		Host: config.DYWebConfig.EmailHost,
	}
	go func() {
		for {
			nowTime := <-ticker.C
			count := 0
			service.Count, count = count, service.Count
			Log.Infof("%s 当前存储item(%d/min) 总共存储item(%d)", nowTime.String(), count, service.SumCount)
			time.Sleep(time.Second / 2)
		}
	}()
	return rpcsupport.ServeRpc(host, service, service1, serviceEmail)
}
