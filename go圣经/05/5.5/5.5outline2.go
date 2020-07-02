package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	//resp , err := http.Get(os.Args[1])
	resp , err := http.Get("http://localhost/5.2findlinks2_2.html")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {

	}
	forEachNode(doc, startElement, endElement)

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

var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attr := ""
		for _, a := range 
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

