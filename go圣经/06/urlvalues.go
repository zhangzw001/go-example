package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang":{"en"}}
	m.Add("item","1")
	m.Add("item","2")

	fmt.Println(m.Get("lang"))

	fmt.Println(m)

	m = nil
	fmt.Println(m.Get("item")) // = Value(nil).Get("item")

	//
	m.Add("item","3")  // panic: assignment to entry in nil map
}
