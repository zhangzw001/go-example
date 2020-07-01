package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack , n.Data) 	//push tag
		fmt.Println(stack )
	}
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		stack = outline(stack, c )
	}
	return stack
}

func outlinecount( n *html.Node) map[string]int {
	outlinemap := make(map[string]int )
	fmt.Println(outlinemap)
	if n.Type == html.ElementNode {
		outlinemap[n.Data]++
	}
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		outlinemap = outlinecount(c)
	}
	return outlinemap
}
func main() {
	//resp, err := http.Get(os.Args[1])
	//resp, err := http.Get("http://localhost/5.2findlinks2_1.html")
	resp, err := http.Get("https://www.baidu.com")
	if err != nil { log.Fatal(err)}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	//outline(nil, doc)
	//sourceStack := []string{}
	//sourceStack = outline(sourceStack,doc)
	//fmt.Println(sourceStack)
	outlineMap := make(map[string]int)
	outlineMap = outlinecount(doc)
	fmt.Println(outlineMap)
}