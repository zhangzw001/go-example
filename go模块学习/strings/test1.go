package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	var a interface{}
	var u = User{
		ID:   1,
		Name: "curry",
	}
	a = u.ID
	fmt.Printf("%T\t%#[1]v\n", a)
	fmt.Sprintln(a)
	//fmt.Println(fmt.Sprintln(a))
	fmt.Println(reflect.TypeOf(fmt.Sprintln(a)))
	var str strings.Builder
	for i := 0 ; i < 1000; i ++ {
		str.WriteString("a")
	}
	fmt.Println(str.String(),str)
}

