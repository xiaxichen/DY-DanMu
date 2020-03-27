package client

import (
	"DY-DanMu/DMconfig/config"
	"DY-DanMu/web/client"
	"DY-DanMu/web/server/_type"
	"DY-DanMu/web/util"
	Log "github.com/sirupsen/logrus"
)

//EmailSend:发送邮件
func EmailSend() (chan _type.EmailSendStruct, error) {
	out := make(chan _type.EmailSendStruct)
	go func() {
		for {
			data := <-out
			var result []string
			err := client.ClientRPC.Call(config.DYWebConfig.SendEmail, data, &result)
			if err != nil {
				err = util.RpcClientShutDownErrorhandler(err)
				if err == nil {
					err = client.ClientRPC.Call(config.SpiderConfig.ItemSaverRpc, data, &result)
				}
			}
			if err != nil {
				Log.Errorf("send email error %v:%v", data, err)
			}
		}
	}()
	return out, nil
}
