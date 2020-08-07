package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)


func main() {

	var w io.Writer
	fmt.Printf("动态类型: %[1]T,动态值: %[1]v\n",w)	//w 此时是空接口, 动态类型为nil, 动态值为nil
	if w == nil {
		fmt.Println("空接口")
	}
	w = os.Stdout
	fmt.Printf("动态类型: %T\n",w)
	if w == nil {
		fmt.Println("空接口")
	}
	w = new(bytes.Buffer)
	fmt.Printf("动态类型: %T\n",w)
	if w == nil {
		fmt.Println("空接口")
	}
	w = nil
	fmt.Printf("动态类型: %T\n",w)
	if w == nil {
		fmt.Println("空接口")
	}

	var t1 io.Writer
	fmt.Println(reflect.TypeOf(t1),reflect.ValueOf(t1))

	var t2 *bytes.Buffer
	fmt.Println(reflect.TypeOf(t2),reflect.ValueOf(t2))

	//然而，如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且panic:
	//考虑到这点，接口类型是非常与众不同的。其它类型要么是安全的可比较类型（如基本类型和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），但是在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic。同样的风险也存在于使用接口作为map的键或者switch的操作数。只能比较你非常确定它们的动态值是可比较类型的接口值。
	//var x interface{} = []int{1, 2, 3}
	//fmt.Println(x == x) // panic: comparing uncomparable type []int
	var s interface{} = "a"
	fmt.Println(s == s )
}