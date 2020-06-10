package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {


	start := time.Now()
	ch2 := make(chan string )

	for _, url := range os.Args[1:] {
		go fetch2(url, ch2)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch2)
	}
	fmt.Printf("time: %.2fs\n",time.Since(start).Seconds())
}

func fetch2(url string , ch chan <- string ) {
	start := time.Now()
	if ! strings.HasPrefix(url, "http") {
		url = strings.Join([]string{"http://",url},"")
	}

	resp , err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes,err :=io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	ch <- fmt.Sprintf("%.2fs  %7d  %s",time.Since(start).Seconds(), nbytes, url)
}
