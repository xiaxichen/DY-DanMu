package main

import (
	"DY-DanMu/DMconfig/config"
	_type "DY-DanMu/spider/DYtype"
	client2 "DY-DanMu/spider/client"
)

func main() {
	itemsChan, err := client2.ItemSaver(config.SpiderConfig.ItemSaverPort)
	if err != nil {
		panic(err)
	}

	client := client2.DyBarrageWebSocketClient{
		Config:      config.SpiderConfig,
		ItemIn:      itemsChan,
		MsgBreakers: _type.CodeBreakershandler{},
	}

	client.Init()
	client.Start()
}
