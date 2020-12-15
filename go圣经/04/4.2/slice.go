package main

import "fmt"

func main() {
	l := []int{1,2,3,4,5}
	fmt.Println(RemoveSliceIndex(l,3))

	l2 := []int{1,2,3,4,5}
	fmt.Println(RemoveSliceIndex2(l2,2))
}

func RemoveSliceIndex(s []int, i int) []int {
	copy(s[i:],s[i+1:])
	return s[:len(s)-1]
}

func RemoveSliceIndex2(s []int, i int) []int {
	return append(s[:i],s[i+1:]... )
}
