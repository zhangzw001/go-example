package main

import (
	"html"
	"os"
)

func main() {
	doc , err := html.Parse(os.Stdin)
}
