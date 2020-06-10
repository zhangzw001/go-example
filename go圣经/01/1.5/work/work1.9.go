// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	list := []string{"www.boqii.com"}

	//for _, url := range os.Args[1:] {
	for _, url := range list {
		if ! strings.HasPrefix(url,"http://")  {
			fmt.Println("url缺少 http://")
			//url="http://"+url
			url = strings.Join([]string{"http://",url},"")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//b, err := io.Copy(os.Stdout,resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
		for k,v := range resp.Header {
			fmt.Printf("%16.16s:\t%s\n",k,v)
		}
		fmt.Println(resp.Status,resp.StatusCode)
	}
}