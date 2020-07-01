package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack , n.Data) 	//push tag
		fmt.Println(stack )
	}
	for c := n.FirstChild; c != nil ; c = c.NextSibling {
		outline(stack, c )
	}
}

func main() {
	//resp, err := http.Get(os.Args[1])
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
	sourceStack := []string{"1"}
	outline(sourceStack,doc)
	fmt.Println(sourceStack)
}