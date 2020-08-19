package main

import (
	"fmt"
	"golang.org/x/net/html"
	"gopl.io/ch5/links"
	"log"
	"net/http"
	"os"
)

func main() {
	//a, err := Extract("http://localhost/5.2findlinks2_1.html")
	//if err != nil { log.Fatal(err)}
	//fmt.Println(a)
	//breadthFirst(crawl, []string{"http://localhost/5.2findlinks2_1.html"})

	worklist := make(chan []string)
	var n int
	n ++
	go func() { worklist <- os.Args[1:]}()

	seen := make(map[string]bool)

	for ; n > 0 ; n -- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n ++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}

}


// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// 限制并发数
var tokens = make(chan struct{}, 10 )
func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil { log.Print(err)}
	return list
}