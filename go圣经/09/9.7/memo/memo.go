package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"memo/memo1"
	"memo/memo2"
	"memo/memo4"
	"net/http"
	"sync"
	"time"
)

func incomingURLs() []string  {
	s := []string{"http://www2.zhihu.com/","http://www.baidu.com","http://www.boqii.com","http://www.qq.com","https://blog.k1s.club","http://www2.zhihu.com/","http://www.baidu.com","http://www.boqii.com","http://www.qq.com","https://blog.k1s.club"}
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

func main() {
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

	m1 := memo1.New(httpGetBody)
	var n1 sync.WaitGroup
	for _,url := range incomingURLs() {
		n1.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m1.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes\n",
				url, time.Since(start),len(value.([]byte)))
			n1.Done()
		}(url)
	}
	n1.Wait()


	start := time.Now()
	m := memo2.New(httpGetBody)
	var n2 sync.WaitGroup
	for _,url := range incomingURLs() {
		n2.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes\n",
				url, time.Since(start),len(value.([]byte)))
			n2.Done()
		}(url)
	}
	n2.Wait()
	fmt.Println(time.Since(start))

	//start := time.Now()
	//m := memo3.New(httpGetBody)
	//var n sync.WaitGroup
	//for _,url := range incomingURLs() {
	//	n.Add(1)
	//	go func(url string) {
	//		start := time.Now()
	//		value, err := m.Get(url)
	//		if err != nil {
	//			log.Print(err)
	//		}
	//		fmt.Printf("%s,%s,%d bytes\n",
	//			url, time.Since(start),len(value.([]byte)))
	//		n.Done()
	//	}(url)
	//}
	//n.Wait()
	//fmt.Println(time.Since(start))

	start4 := time.Now()
	m4 := memo4.New(httpGetBody)
	var n sync.WaitGroup
	for _,url := range incomingURLs() {
		n.Add(1)

		m4.Get(url)
		go func(url string) {
			start := time.Now()
			value, err := m4.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes\n",
				url, time.Since(start),len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
	fmt.Println(time.Since(start4))
}
