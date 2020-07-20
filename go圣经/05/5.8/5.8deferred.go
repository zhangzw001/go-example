package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var m = make(map[string]int)

func main() {
	m["aa"]=123
	fmt.Println(lookup("aa"))
}

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}