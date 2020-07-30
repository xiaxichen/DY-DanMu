package client

import (
	"DY-DanMu/DMconfig/config"
	"DY-DanMu/persistServer/item"
	_type "DY-DanMu/spider/DYtype"
	"fmt"
	"github.com/gorilla/websocket"
	Log "github.com/sirupsen/logrus"
	"log"
	"time"
)

/* DyBarrageWebSocketClientInterface: 通过斗鱼open文档
https://open.douyu.com/source/api/63 进行弹幕服务器连接抓取
*/
type DyBarrageWebSocketClientInterface interface {
	Start()
	Stop()
	save(item map[string]string)
	send(msg string) error
	Init()
	getOnMsg()
	runForever()
	login()
	joinGroup()
	startHeartbeat()
	logout()
	onError(err error)
}

//DyBarrageWebSocketClient:斗鱼弹幕服务器连接端
type DyBarrageWebSocketClient struct {
	ws                      *websocket.Conn
	MsgBreakers             _type.CodeBreakershandler
	__should_stop_heartbeat bool
	Config                  *config.DMconfig
	sentry                  chan int
	ItemIn                  chan item.Item
}

func (d *DyBarrageWebSocketClient) Init() {
	dial, _, err := websocket.DefaultDialer.Dial(d.Config.Url, nil)
	if err != nil {
		panic(err)
	}
	d.ws = dial
	d.sentry = make(chan int)
	d.__should_stop_heartbeat = false
}

//Start:启动
func (d *DyBarrageWebSocketClient) Start() {
	d.runForever()
}

//Stop:停止
func (d *DyBarrageWebSocketClient) Stop() {
	d.ws.Close()
	d.logout()
}

//send:发送编码过的数据到socket服务器
func (d *DyBarrageWebSocketClient) send(msg string) error {
	err := d.ws.WriteMessage(websocket.TextMessage, d.MsgBreakers.Encode(msg))
	return err
}

//save:保存数据
func (d *DyBarrageWebSocketClient) save(res map[string]string) {
	items := item.Item{
		Rid:     res["rid"],
		Id:      res["cid"],
		Payload: res,
	}
	d.ItemIn <- items
}

//getOnMsg:从DY服务器端获取弹幕赫尔状态进行解析
func (d *DyBarrageWebSocketClient) getOnMsg() {
	for {
		status, message, err := d.ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			d.onError(err)
			continue
		}
		switch {
		case status == 0:
			fmt.Println(status)
		case (status == 1):
			fmt.Println(message)
		case (status == 2):
			messages := d.MsgBreakers.GetChatMessages(message)
			for _, msg := range messages {
				go d.save(msg)
			}
		case (status == 8):
			d.Stop()
		default:
			fmt.Println(status)
		}
	}
}

//runForever:程序入口
func (d *DyBarrageWebSocketClient) runForever() {
	d.login()
	d.joinGroup()
	go d.startHeartbeat()
	go d.getOnMsg()
	<-d.sentry
}

//login:发送登录信息
func (d *DyBarrageWebSocketClient) login() {
	err := d.send(fmt.Sprintf(d.Config.LoginMsg, d.Config.Rid, "61609154", "61609154"))
	if err != nil {
		panic(err)
	}
}

//joinGroup:加入服务器端群组中
func (d *DyBarrageWebSocketClient) joinGroup() {
	err := d.ws.WriteMessage(websocket.TextMessage, d.MsgBreakers.Encode(
		fmt.Sprintf(d.Config.LoginJoinGroup, d.Config.Rid),
	))
	if err != nil {
		panic(err)
	}
}

//startHeartbeat:保持与服务端的心跳每45秒发送一次
func (d *DyBarrageWebSocketClient) startHeartbeat() {
	heartbeat_msg := "type@=mrkl/"
	heartbeat_msg_byte := d.MsgBreakers.Encode(heartbeat_msg)
	for {
		err := d.ws.WriteMessage(websocket.TextMessage, heartbeat_msg_byte)
		for i := 0; i < 90; i++ {
			time.Sleep(time.Millisecond * 500)
			if err != nil {
				log.Fatal(err)
			}
			if d.__should_stop_heartbeat {
				d.ws.Close()
				d.sentry <- 1
				return
			}
		}
	}

}

//logout:登出服务器
func (d *DyBarrageWebSocketClient) logout() {
	logoutMsg := "type@=logout/"
	logoutMsgByte := d.MsgBreakers.Encode(logoutMsg)
	d.__should_stop_heartbeat = true
	log.Println(logoutMsgByte)
}

//onError:处理异常
func (d *DyBarrageWebSocketClient) onError(err error) {
	Log.Warnf("socker error! %s", err)
	d.ws.Close()
	dial, _, err := websocket.DefaultDialer.Dial(d.Config.Url, nil)
	if err != nil {
		panic(err)
	}
	d.ws = dial
	d.login()
	d.joinGroup()
}
