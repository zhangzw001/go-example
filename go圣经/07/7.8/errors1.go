package main

import (
	"fmt"
)

type p struct {
	//s string
}
func main() {
	//fmt.Println(errors.New("EOF") == errors.New("EOF")) //false
	p3 := &p{}
	p4 := &p{}
	//fmt.Println(p3,p4)
	fmt.Println(p3 == p4 )
	//fmt.Println(&p3 == &p4 )
	//fmt.Println(*p3 == *p4 )
}
