package storageMail

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 95
}


// email sender configuration
// note: never put password in source code!
const sender = "bq_zabbix@boqii.com"
const password = "boqii@3124"
const hostname = "smtp.boqii.com"


const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`


func CheckQuota(username string) {
	used := bytesInUse(username)
	fmt.Println(used)
	const quota = 1000000000 //1GB
	percent := 100 * used / quota
	if percent < 90 {
		log.Printf("disk < %d% \n",percent)
		return
	}else {
		log.Printf("disk >= %d% \n",percent)
	}
	msg := fmt.Sprintf(template,used,percent)
	auth := smtp.PlainAuth("",sender, password,hostname)
	err := smtp.SendMail(hostname+":25",auth,sender, []string{username},[]byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s",username,err)
	}else {
		log.Printf("smtp.SendMail(%s) successed: %s",username,err)
	}
}