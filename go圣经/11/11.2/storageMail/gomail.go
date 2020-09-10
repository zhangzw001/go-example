package storageMail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type mail struct {
	serverHost string
	serverPort int
	fromUser   string
	fromPass   string
}

func (m *mail )init() {
	fmt.Println("init")
	*m = mail{
		serverHost: "smtp.boqii.com",
		serverPort: 465,
		fromUser: "bq_zabbix",
		fromPass: "boqii@3124",
	}
}

func GomailLearn1() {
	var m1 mail
	m1.init()
	fmt.Println(m1)
	d := gomail.NewDialer(m1.serverHost,m1.serverPort,m1.fromUser,m1.fromPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From","bq_zabbix@boqii.com")
	m.SetHeader("To","zhangzw@boqii.com")
	m.SetAddressHeader("Cc","zhangzw@boqii.com","zhangzw")
	m.SetBody("text/html","hello <b>zhangzw</b>  <a href=\"www.baidu.com\">www.baidu.com</a>")

}
