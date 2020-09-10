package main

import (
	"word/storageMail"
)

func main() {
	//seed := time.Now().UTC().UnixNano()
	//fmt.Printf("Random seed: %d",seed)
	//r := rand.New(rand.NewSource(seed))
	//fmt.Println(0x1000)
	//fmt.Println(0x0FFF)
	//fmt.Println(r.Intn(0x0999-0x1000)+0x999)
	//fmt.Println(word.IsPalindrome2("été"))
	storageMail.CheckQuota("zhangzw@boqii.com")


}