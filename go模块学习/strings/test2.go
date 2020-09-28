package main

import (
	"fmt"
	"reflect"
)

const s = "Go101.org"
// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128
var b1 byte = 1 << len(s[:])
//
func main() {
	fmt.Println(reflect.TypeOf(b1),b1 )
	println(a, b)
}
