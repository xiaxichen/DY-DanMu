package main

import (
	"DY-DanMu/DMconfig/config"
	_type "DY-DanMu/spider/DYtype"
	client2 "DY-DanMu/spider/client"
	Log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	saverStruct := client2.ItemSaverStruct{
		ItemCountAll: 0,
		ItemCountMin: 0,
	}
	itemsChan, err := saverStruct.ItemSaver(config.SpiderConfig.ItemSaverPort)
	if err != nil {
		panic(err)
	}
	emailSendChan, err := client2.EmailSend()
	if err != nil {
		panic(err)
	}
	client := client2.DyBarrageWebSocketClient{
		Config: config.SpiderConfig,
		ItemIn: itemsChan,
		MsgBreakers: _type.CodeBreakershandler{
			IsLive:        false,
			EmailSendChan: emailSendChan,
		},
	}
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			nowTime := <-ticker.C
			count := 0
			saverStruct.ItemCountMin, count = count, saverStruct.ItemCountMin
			Log.Infof("%s 当前存储item(%d/min) 总共存储item(%d)", nowTime.String(), count, saverStruct.ItemCountAll)
			time.Sleep(time.Second / 2)
		}
	}()
	client.Init()
	client.Start()
}
