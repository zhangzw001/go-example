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
func main() {
	//程序不会停止,即时已经完成爬取
	//能够停止
	crawl_one()
}

func web_crawl_one(url string) map[string][]string {
	fmt.Println(url)
	list, err := web_Extract(url)
	if err != nil { log.Println(err)}
	return list
}
func web_crawl_two(url string) map[string][]string {
	fmt.Println(url)
	//fmt.Printf("当前深度: %v\n",depthFirst)
	if depthFirst >= depths {
		return nil
	}
	depthFirst++
	tokens <- struct{}{}
	fmt.Printf("tokens : %v\n",tokens)
	list, err := web_Extract(url)
	//fmt.Println(list)
	defer func() { <-tokens }()
	if err != nil { log.Println(err)}
	return list
}

func web_Extract(url string ) (map[string][]string , error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil { return nil, err }
	if resp.StatusCode != http.StatusOK {return nil , fmt.Errorf("getting %s: %s",url,resp.Status)}

	doc, err := html.Parse(resp.Body)
	if err != nil { return nil , fmt.Errorf("parsing %s as HTML: %v", url, err )}

	links := make(map[string][]string)
	visitNode := func(n *html.Node) {
		// html.ElementNode 子节点的html属性,
		//fmt.Println(n.Attr)
		if n.Type == html.ElementNode {
			switch n.Data {
			case "img":
					for _, a := range n.Attr {
						if a.Key != "src" {
							continue
						}
						link, err := resp.Request.URL.Parse(a.Val)
						if err != nil {
							continue
						} //失败的url 忽略
						links[n.Data] = append(links[n.Data], link.String())
					}
			case "a":
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					} //失败的url 忽略
					links[n.Data] = append(links[n.Data], link.String())
				}
			default:
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
	for c := n.FirstChild ; c != nil ; c = c.NextSibling {
		web_forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}


//
func crawl_one() {
	worklist := make(chan []string)

	//
	var n int
	n ++
	go func() {
		list := []string{
			"http://www.boqii.com/baike",
			//"http://gopl.io/",
			//"http://gopl.io/",
			//"https://golang.org/help/",
			//"https://golang.org/doc/",
			//"https://golang.org/blog/",
		}
		worklist <- list
	}()

	//并发爬取
	seen := make(map[string]bool)

	for ; n > 0 ; n -- {
		//阻塞, 等到worklist 写入数据
		list := <-worklist
		//fmt.Printf("list: %v\n",list)
		//fmt.Printf("seen: %v\n",seen)
		for _, link := range list {
			// 只要发现一个没有爬过的连接, 就n++, 就继续循环一次, 从chan worklist 继续读取新url
			if !seen[link ] {
				seen[link] = true
				n ++
				go func(link string) {
					//worklist <- web_crawl_one(link)
					worklist <- web_crawl_two(link)["a"]
				}(link)
			}
		}
	}
}