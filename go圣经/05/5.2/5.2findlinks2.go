package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url  := os.Args[1]
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil { log.Fatal(err)}
	if resp.StatusCode != http.StatusOK {
		if err := resp.Body.Close(); err != nil { log.Fatal(err )}
		log.Fatalf("http.get error : %s\n",resp.Status)
	}

	result,err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatal(err )}

	//这里ioutil.ReadAll 返回的是[]byte, 开始不知道bytes.NewReader
	//resultString :=fmt.Sprintf("%s\n",result)
	//doc , err := html.Parse(strings.NewReader(resultString))
	//
	doc , err := html.Parse(bytes.NewReader(result))
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err); err != nil { log.Fatal(err)}
		os.Exit(1)
	}

	//for _, link := range visit2(nil, doc ) {
	//	fmt.Println(link)
	//}
	for  range visit2(nil, doc ) {
	}
}

// visit appends to links each link found in n and returns the result.
func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(links, c )
		fmt.Println("-----------")
		fmt.Println(c)
	}
	return links
}