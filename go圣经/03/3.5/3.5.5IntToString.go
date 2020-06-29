package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(reflect.TypeOf(y))

	fmt.Println(reflect.TypeOf(strconv.Itoa(x)))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(reflect.TypeOf(s),s)


	m , _ := strconv.Atoi("999")
	fmt.Println(reflect.TypeOf(m),m)
	n , _ := strconv.ParseInt("123",10,0)
	fmt.Println(reflect.TypeOf(n),n)
}
