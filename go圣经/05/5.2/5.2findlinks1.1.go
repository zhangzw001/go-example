package main

import (
	"fmt"
	"golang.org/x/net/html"
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

	//result,err := ioutil.ReadAll(resp.Body)
	//if err != nil { log.Fatal(err )}
	//doc , err := html.Parse(bytes.NewReader(result))

	doc, err := html.Parse(resp.Body)
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err); err != nil { log.Fatal(err)}
		os.Exit(1)
	}

	//for _, link := range visit3(nil, doc ) {
	//	fmt.Println(link)
	//}

	//for _, link := range visittext(nil, doc ) {
	//	fmt.Println(link)
	//}
	map1 := make(map[string][]string)
	for k, link := range visit4(map1, doc ) {
		fmt.Println(k,link)
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
	}
	return links
}

// visit appends to links each link found in n and returns the result.
//练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
func visit3(links []string, n *html.Node) []string {
	//递归出口
	if n == nil { return links }
	links = visit3(links,n.FirstChild)
	links = visit3(links,n.NextSibling)

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//}
	return links
}



//练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
func visittext(text []string , n *html.Node) []string {
	if n.Type == html.TextNode {
		text = append(text, n.Data)
	}
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		text = visittext(text,c)
	}
	return text
}

//练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
func visit4(links map[string][]string, n *html.Node) map[string][]string {

	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "style" ){
		for _, a := range n.Attr {
			switch a.Key {
			case "href":
				links["href"] = append(links["href"],a.Val)
			case "src":
				links["img"] = append(links["img"],a.Val)
			case "type":
				links["style"] = append(links["style"],a.Val)
			default:
				return links
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit4(links, c )
	}
	return links
}
