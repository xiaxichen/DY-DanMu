package main

import (
	"DY-DanMu/DMconfig/config"
	"DY-DanMu/dbMysql/mysql"
	"DY-DanMu/persistServer/rpcsupport"
	"DY-DanMu/persistServer/server"
	"github.com/olivere/elastic/v7"
	Log "github.com/sirupsen/logrus"
)

func main() {
	Log.Fatal(serverRpc(config.SpiderConfig.ItemSaverPort, config.SpiderConfig.ElasticIndex))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	service := &server.ItemSaverService{
		Client: client,
		Conn:   mysql.DBConn(),
		Index:  index,
		Count:  0,
	}
	service1 := &server.SelectMiddlerWare{
		Client: client,
		Conn:   mysql.DBConn(),
		Index:  config.DYWebConfig.ElasticIndex,
		Count:  0,
	}

	return rpcsupport.ServeRpc(host, service, service1)
}
