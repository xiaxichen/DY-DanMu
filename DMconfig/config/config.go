package config

import "os"

type DMconfig struct {
	Rid            string
	LoginMsg       string
	LoginJoinGroup string
	Url            string
	ItemSaverPort  string
	ElasticIndex   string
	ItemSaverRpc   string
}

type DYWebSeverConfig struct {
	Host                  string
	UserSearch            string
	BarrageAll            string
	SearchAllField        string
	ElasticIndex          string
	IndexBarrageCount     string
	StatisticsBarrage     string
	StatisticsUserBarrage string
	SendEmail             string
	EmailHost             string
	EmailUser             string
	EmailPwd              string
}

var SpiderConfig *DMconfig
var DYWebConfig *DYWebSeverConfig

func init() {
	SpiderConfig = &DMconfig{
		Rid:            "48699",
		LoginMsg:       "type@=loginreq/room_id@=%s/dfl@=sn@A=105@Sss@A=1/username@=%s/uid@=%s/ver@=20190610/aver@=218101901/ct@=0/",
		LoginJoinGroup: "type@=joingroup/rid@=%s/gid@=-9999/",
		Url:            "wss://danmuproxy.douyu.com:8506/",
		ItemSaverPort:  ":5100",
		ElasticIndex:   "dou_yu_barrage",
		ItemSaverRpc:   "ItemSaverService.Save",
	}
	emailUser := os.Getenv("EMAILUSER")
	emailPwd := os.Getenv("EMAILPWD")
	DYWebConfig = &DYWebSeverConfig{
		Host:                  ":5100",
		UserSearch:            "SelectMiddlerWare.UserQuery",
		BarrageAll:            "SelectMiddlerWare.BarrageAll",
		SearchAllField:        "SelectMiddlerWare.SearchFieldAll",
		IndexBarrageCount:     "SelectMiddlerWare.BarrageCount",
		StatisticsBarrage:     "SelectMiddlerWare.StatisticsBarrageForTime",
		StatisticsUserBarrage: "SelectMiddlerWare.StatisticsUserBarrageForTime",
		SendEmail:             "EmialSendSever.SendToMail",
		ElasticIndex:          "dou_yu_barrage",
		EmailHost:             "smtp.163.com:25",
		EmailUser:             emailUser,
		EmailPwd:              emailPwd,
	}
}
