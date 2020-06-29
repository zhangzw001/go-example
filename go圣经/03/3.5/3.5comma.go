package main

import (
	"fmt"
	comma "github.com/zhangzw001/go-example/go圣经/03/3.5/3.5comma"
)

func main() {

	s := "-12345678900000.11111"
	fmt.Println(comma.Comma(s))
	fmt.Println(comma.CommaBuffer(s))
}
