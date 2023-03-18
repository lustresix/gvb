package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

func send(name, subject, body string) error {
	return sendEmail(
		viper.GetString("email.user"),
		viper.GetString("email.password"),
		viper.GetString("email.host"),
		viper.GetInt("email.port"),
		name,
		viper.GetString("email.user"),
		subject,
		body,
	)
}

func sendEmail(userName, authCode, host string, port int, mailTo, sandName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sandName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}
