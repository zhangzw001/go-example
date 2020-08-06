package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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



}