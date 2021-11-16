package common

import (
	"2022/ginseckill/config"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmailCode(toEmail,msg string) (err error) {
	ec := config.Conf.Email
	fromEmail := ec.FromEmail
	subject := ec.EmailSubject
	smtpAddr := ec.SmtpAddr
	smtpPort,_ := strconv.Atoi(ec.SmtpPort)
	smtpPass := ec.SmtpPass
	mailHeader := map[string][]string {
		"From":{fromEmail},
		"To":{toEmail},
		"Subject":{subject},
	}
	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	m.SetBody("text/html","尊敬的用户<br>您的注册验证码为:"+msg+"<br>欢迎您!")
	d := gomail.NewDialer(smtpAddr,smtpPort,fromEmail,smtpPass)
	err = d.DialAndSend(m)
	return
}