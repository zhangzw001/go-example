package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	outline3()

}

func outline3() {
	var depth int
	//resp , err := http.Get(os.Args[1])
	resp , err := http.Get("http://localhost/5.2findlinks2_2.html")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {

	}
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			attr := ""
			for _, a := range n.Attr {
				attr = fmt.Sprintf(" %s=\"%s\" ",a.Key,a.Val)
			}
			fmt.Printf("%*s<%s%s", depth*2, "", n.Data,attr)
			depth++
		}
		if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
			fmt.Printf("/>\n")
		} else if n.Type == html.ElementNode {
			fmt.Printf(">\n")
		}

		if n.Type == html.TextNode {
			fmt.Printf("%*s %s\n", depth*2,"",n.Data)
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
			depth--
			//fmt.Printf("\n")
			return
		}
		if n.Type == html.ElementNode {
			depth--

			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	forEachNode3(doc, startElement, endElement)

}
// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode3(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode3(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}


