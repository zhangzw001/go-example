package main

import (
	"fmt"
	"strings"
)

/*
练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
*/
func expand(s string, f func(string) string) string {
	str := f("foo")
	s = strings.Replace(s, "foo", str, -1)
	return s
}
func expand2(s string) string {
	return s + "-test"
}
func main() {
	s1 := "foo"
	fmt.Println(expand(s1 ,expand2))
}
