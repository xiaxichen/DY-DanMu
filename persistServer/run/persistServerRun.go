package main

import (
	"DY-DanMu/DMconfig/config"
	"DY-DanMu/dbMysql/mysql"
	"DY-DanMu/persistServer/rpcsupport"
	"DY-DanMu/persistServer/server"
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
	service := &server.ItemSaverService{
		Client:   client,
		Conn:     mysql.DBConn(),
		Index:    index,
		Count:    0,
		SumCount: 0,
	}
	service1 := &server.SelectMiddlerWare{
		Client: client,
		Conn:   mysql.DBConn(),
		Index:  config.DYWebConfig.ElasticIndex,
		Count:  0,
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
