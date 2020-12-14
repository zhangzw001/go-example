package main

import "fmt"

func main() {
	l := []int{1,2,3,4}
	l2 := append(l[:2],l[3:]... )
	fmt.Println(l2,l)
}
