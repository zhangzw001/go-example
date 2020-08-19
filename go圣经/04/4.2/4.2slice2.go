package main

import "fmt"

func main() {
	a := []int{0,1,2,3,4,5,6,7,8,9}
	s6 := a[2:5]
	s6[2] = 888
	s7 := s6[2:7]
	s7[2] = 999
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Println(a)
	fmt.Print(cap(s6),cap(s7))
}
