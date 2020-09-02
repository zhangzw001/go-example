package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"memo/memo1"
	"net/http"
	"sync"
	"testing"
	"time"
)

func incomingURLs() []string  {
	s := []string{"http://www.baidu.com"}
	return s
}

func httpGetBody(url string) (interface{}, error) {
	resp,err := http.Get(url)
	if err != nil {
		return nil , err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


func TestMemo1(t *testing.T) {
	//m := memo1.New(httpGetBody)
	//for _,url := range incomingURLs() {
	//	start := time.Now()
	//	value, err := m.Get(url)
	//	if err != nil {
	//		log.Print(err)
	//	}
	//	fmt.Printf("%s, %s, %d bytes\n",
	//		url, time.Since(start), len(value.([]byte)))
	//}

	m := memo1.New(httpGetBody)
	var n sync.WaitGroup
	for _,url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes\n",
				url, time.Since(start),len(value.([]byte)))
			n.Done()
		}(url)
	}

}
