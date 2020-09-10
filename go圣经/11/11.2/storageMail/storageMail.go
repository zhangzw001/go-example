package storageMail

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 950000000
}


// email sender configuration
// note: never put password in source code!
const sender = "tideways_bq5@163.com"
const password = "KWAAJNELLTSHYAIV"
const hostname = "smtp.163.com"


const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`


func CheckQuota(username string) {
	used := bytesInUse(username)
	fmt.Println(used)
	const quota = 1000000000 //1GB
	percent := 100 * used / quota
	log.Printf("disk is %d% \n",percent)
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template,used,percent)
	auth := smtp.PlainAuth("",sender, password,hostname)
	err := smtp.SendMail(hostname+":25",auth, sender, []string{username},[]byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s",username,err)
	}else {
		log.Printf("smtp.SendMail(%s) successed: %s",username,err)
	}

}