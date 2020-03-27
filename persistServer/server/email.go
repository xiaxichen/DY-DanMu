package server

import (
	"DY-DanMu/web/server/_type"
	Log "github.com/sirupsen/logrus"
	"net/smtp"
)

type EmialSendSever struct {
	Auth smtp.Auth
	Host string
}

//SendToMail:发送邮件
func (e EmialSendSever) SendToMail(EST _type.EmailSendStruct, result *[]string) error {
	var content_type string
	if EST.MailType == _type.EMAILHTMLTYPE {
		content_type = "Content-Type: text/" + EST.MailType + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	for _, to := range EST.To {
		msg := []byte("To: " + to + "\r\nFrom: " + EST.UserName + "\r\nSubject: " + "夜吹直播开始了!" + "\r\n" +
			content_type + "\r\n\r\n" + EST.Body)
		//sendTo := strings.Split(to, ";")
		sendTo := []string{"xiaxichen1@163.com", "337094778@qq.com"}
		err := smtp.SendMail(e.Host, e.Auth, EST.UserName, sendTo, msg)
		if err != nil {
			Log.Error(err)
			*result = append(*result, "err")
		} else {
			*result = append(*result, "ok")
		}
	}
	return nil
}
