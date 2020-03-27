package server

import (
	config2 "DY-DanMu/DMconfig/config"
	"DY-DanMu/dbMysql/config"
	"DY-DanMu/lib"
	"DY-DanMu/persistServer/item"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	Log "github.com/sirupsen/logrus"
	"time"
)

type ItemSaverService struct {
	Client   *elastic.Client
	Conn     *sql.DB
	Index    string
	Count    int
	SumCount int
}

// Save:存储数据
func (s *ItemSaverService) Save(item item.Item, result *string) error {
	err, err1 := Save(s.Client, s.Conn, item, s.Index)
	if err == nil && err1 == nil {
		s.Count++
		s.SumCount++
		Log.Debug("RPC Count ItemSaver: %d", s.Count)
		*result = "ok"
	} else {
		Log.Errorf("Error Saving item %v: %v %v", item, err)
	}
	return err
}

// Save: 处理/保存数据到Es和Mysql
func Save(client *elastic.Client, conn *sql.DB, item item.Item, index string) (err error, err1 error) {
	Payload := lib.CheckIt(item.Payload)
	if Payload != nil {
		Scst, _ := Payload["cst"].(string)
		if Scst == "" {
			Payload["cst"] = time.Now().UnixNano() / 1e6
		} else {
			Payload["cst"] = lib.TrunType(Scst)
		}
		Sbl, _ := Payload["bl"].(string)
		Payload["bl"] = lib.TrunType(Sbl)
		Slevel, _ := Payload["level"].(string)
		Payload["level"] = lib.TrunType(Slevel)
		Surlev, _ := Payload["urlev"].(string)
		Payload["urlev"] = lib.TrunType(Surlev)
		item.Payload = Payload
		indexService := client.Index().Index(index).BodyJson(item)
		if item.Id != "" {
			indexService.Id(item.Id)
		}
		_, err = indexService.Do(context.Background())
		prepare, err1 := conn.Prepare(fmt.Sprintf("insert ignore into %s(`cid`,`level`,`bl`,`bnn`,`brid`,`col`,`cst`,`dms`,`ifs`,`esIndex`,`hc`,`urlev`,`type`,`sahf`,`lk`,`fl`,`el`,`ct`,`txt`,`uid`,`nn`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", config.MysqlDBName+"."+config.MysqlTableName))
		if err != nil && err1 != nil {
			return err, err1
		}
		defer prepare.Close()
		sqlResult, err1 := prepare.Exec(Payload["cid"], Payload["level"], Payload["bl"], Payload["bnn"], Payload["brid"],
			Payload["col"], Payload["cst"], Payload["dms"], Payload["ifs"], config2.SpiderConfig.ElasticIndex, Payload["hc"],
			Payload["urlev"], Payload["type"], Payload["sahf"], Payload["lk"], Payload["fl"], Payload["el"], Payload["ct"], Payload["txt"], Payload["uid"], Payload["nn"])
		if err != nil && err1 != nil {
			return err, err1
		}
		affected, err := sqlResult.RowsAffected()
		if affected <= 0 && err == nil {
			return nil, errors.New("insert data is exit")
		}
		return nil, nil
	} else {
		return errors.New("intertface to Map string error!"), nil
	}
}
