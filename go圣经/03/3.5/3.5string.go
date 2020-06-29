package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`


func main() {
	s := "hello world"
	fmt.Printf("%s %d \n%s\n",s[0:len(s)],len(s),s[0:7])


	fmt.Println("goodbye " + s[6:])
	fmt.Println(strings.Join([]string{"goodbye",s[6:]}," "))

	fmt.Println(GoUsage)

	s1 := "api_goo"
	fmt.Println(HasPrefix(s1,"api"))
	fmt.Println(HasSuffix(s1,"goo"))
	fmt.Println(Contains(s1,"pi"))

	s2 := "Hello, 世界"
	fmt.Println(len(s2))                    // "13"
	fmt.Println(utf8.RuneCountInString(s2)) // "9"

	n := 0
	for i, r := range s2 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		n ++
	}
	fmt.Printf("s2的字符个数是: %d, s2的len()是: %d\n",n,len(s2))
	fmt.Println(utf8.RuneCountInString(s2))


	s3 := "abc"
	b := []byte(s3)
	s4 := string(b)
	fmt.Println(s3,b,s4)
	fmt.Printf("%p %p %p\n",s3,b,s4)
	fmt.Println(reflect.TypeOf(b),reflect.TypeOf(s4))


}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}