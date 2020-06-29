package main

import (
	"fmt"
)

func main() {
	s1 := "abcd"
	s2 := "badc"
	s3 := "abce"
	fmt.Println(DiffString(s1,s2))
	fmt.Println(DiffString(s1,s3))
}


func DiffString(s1,s2 string ) bool {
	n1 , n2 := len(s1),len(s2)
	flag := true
	if n1 != n2 {
		flag = false
		return flag
	}

	m := make(map[rune]int,0)
	n := make(map[rune]int,0)

	for _,i := range s1  {
		m[i]++
	}
	for _,i := range s2 {
		n[i]++
	}

	for i,_ := range m {
		if m[i] != n[i] {
			flag = false
		}

	}

	return flag
}

