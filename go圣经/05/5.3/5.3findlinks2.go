package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links , err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n",err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string ) ([]string , error ) {
	resp , err := http.Get(url )
	if err != nil {
		return nil , err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("gettings %s : %s ")
	}
}