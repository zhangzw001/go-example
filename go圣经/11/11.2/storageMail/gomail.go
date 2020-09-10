package storageMail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type mail struct {
	serverHost 	string
	serverPort 	int
	fromUser   	string
	fromPass   	string
	msg 		*gomail.Message
}

func (m *mail )init() {
	fmt.Println("init")
	*m = mail{
		serverHost: "",
		serverPort: 1,
		fromUser: "",
		fromPass: "",
		msg: gomail.NewMessage(),
	}
	m.msg.SetHeader("From","bq_zabbix@boqii.com")
	m.msg.SetHeader("To","zhangzw@boqii.com")
	m.msg.SetAddressHeader("Cc","zhangzw@boqii.com","zhangzw")
	m.msg.SetBody("text/html","hello <b>zhangzw</b>  <a href=\"www.baidu.com\">www.baidu.com</a>")
	m.msg.SetHeader("Subject", "Hello!")

}

func GomailLearn1() {
	var m1 mail
	m1.init()

	d := gomail.NewDialer(m1.serverHost,m1.serverPort,m1.fromUser,m1.fromPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m1.msg); err != nil {
		panic(err)
	}

}
