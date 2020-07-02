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
//练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img>
//</img>）。编写测试，验证程序输出的格式正确。（详见11章）
//优化了script标签问题
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attr := ""
		for _, a := range n.Attr {
			attr = fmt.Sprintf(" %s=\"%s\" ",a.Key,a.Val)
			//attr += " " + a.Key + "=" + "\"" + a.Val + "\" "
			//fmt.Println(a.Key, a.Val, n.Data, attr)

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
func endElement(n *html.Node) {
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

