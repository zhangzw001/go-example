package main

import (
	"fmt"
	"reflect"
	"runtime"
)

const s = "Go101.org"
// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

// 这里是 常量位移表达式
var a byte = 1 << len(s) / 128
//var a1 byte = 1 << len(s)
// 这里是 非常量位移表达式, 所以会位移后转换为 byte类型 然后再做除操作,所以是 0 / 128
var b byte = 1 << len(s[:]) / 128
var b1 byte = 1 << len(s[:])
//
func main() {
	fmt.Println(reflect.TypeOf(b1),b1 )
	//fmt.Println(reflect.TypeOf(a1),a1 )
	println(a, b)

	arr := []int{1,2,3}
	narr := []*int{}
	for i,_ := range arr {
		narr = append(narr,&arr[i])
	}
	//for _, v := range arr {
	//	narr = append(narr,&v)
	//}
	runtime.m
	for _, v := range narr {
		fmt.Println(*v)
	}
}
