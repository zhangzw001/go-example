package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

/*
   练习 8.6： 为并发爬虫增加深度限制。
   也就是说，如果用户设置了depth=3，
   那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
*/

var depths int = 3
var depthFirst int = 0
var tokens = make(chan struct{}, 20)

func web_crawl_two(url string) []string {
	fmt.Println(url)
	if depthFirst >= depths {
		return nil
	}
	depthFirst++
	tokens <- struct{}{}
	fmt.Printf("tokens : %v\n",tokens)
	list, err := web_Extract(url)
	defer func() { <-tokens }()
	if err != nil { log.Println(err)}
	return list
}

func web_Extract(url string ) ([]string , error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil { return nil, err }
	if resp.StatusCode != http.StatusOK {return nil , fmt.Errorf("getting %s: %s",url,resp.Status)}

	doc, err := html.Parse(resp.Body)
	if err != nil { return nil , fmt.Errorf("parsing %s as HTML: %v", url, err )}

	var links []string
	visitNode := func(n *html.Node) {
		// html.ElementNode 就是子节点的html属性, 
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link , err := resp.Request.URL.Parse(a.Val)
				if err != nil { continue} //失败的url 忽略
				links = append(links, link.String())
			}
		}
	}

	web_forEachNode(doc, visitNode, nil )
	return links, nil
}

func web_forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild ;
}
