package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://www.baidudd.com"
	if err := WaitForServer(url) ; err != nil {
		log.Fatalf("Site is down: %v\n",err)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	// 设置前缀
	log.SetPrefix(">>> ")
	// 替换默认的时间戳
	log.SetFlags(0)
	// log.PrintX 不会退出,  log.FatalX 会调用os.Exit
	// log包中的所有函数会为没有换行符的字符串增加换行符。
	for tries :=0 ; time.Now().Before(deadline) ; tries ++ {
		log.Println("start http.Head")
		_, err := http.Head(url)
		log.Println("end http.Head")
		if err == nil {
			return nil //success
		}
		log.Printf("server not responding (%s); retrying ... ", err)
		a := time.Second << uint(tries)
		fmt.Println(a,tries)
		time.Sleep(a) //exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s",url, timeout)
}