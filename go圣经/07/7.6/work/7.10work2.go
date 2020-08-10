package main

import (
	"fmt"
	"os"
	"sort"
)

type stringSlice []byte

func (s stringSlice) Len () int {return len(s)}
func (s stringSlice) Swap (i,j int) {
	s[i], s[j] = s[j], s[i]
}
func (s stringSlice) Less(i,j int) bool {
	return s[i] < s[j]
}

// IsPalindrome 可以接收任何实现了接口的类型
func IsPalindrome2 (s sort.Interface) bool {
	for i,j:=0,s.Len()-1;i<j; i,j = i+1,j-1 {
		fmt.Println(s.Less(i,j),s.Less(j,i))
		if !(!s.Less(i,j) && !s.Less(j,i)) {
			return false
		}
	}
	return true
}

func main() {
	ss := stringSlice([]byte("abcdba"))
	fmt.Fprintln(os.Stdout, IsPalindrome2(ss))
}
