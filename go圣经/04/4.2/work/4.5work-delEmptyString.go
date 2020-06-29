package main

import "fmt"

//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。


func delEmptyString(s []string ) []string {

	len := len(s)
	n := 0
	for i, _ := range s {
		index := i+1
		if index >= len {
			s[n] = s[i]
			break
		}
		if s[index] != s[i] {
			s[n] = s[i]
			n++
		}
	}
	return s[:n+1]
}


func main() {
	s := []string{"1","1","2","2","2","3"}
	fmt.Println(delEmptyString(s))
	fmt.Println(s)


}