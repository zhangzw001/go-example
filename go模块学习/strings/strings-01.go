package main

import (
	"fmt"
	"strings"
)

func main() {


	// 01
	string1 := "abcd222xxx dfadf hello world! good bye !"

	a := strings.Index(string1, "good bye")
	fmt.Println(a,string1[a:],string1[:a])
}
