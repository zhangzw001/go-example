package main

import (
	"fmt"
	"log"
	"os"
)

func f() {}

var cwd string

func init() {
	//cwd, err := os.Getwd() // compile error: unused: cwd
	var err error
	cwd, err = os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {
	////作用域
	//x := "hi"
	//fmt.Printf("%p,%v\n",x,&x)
	//for _, x := range x {
	//	//fmt.Printf("循环内1:%c,%p,%v\n",x,x,&x)
	//	x := x + 'A' - 'a'
	//	//fmt.Printf("循环内2:%c,%p,%v\n",x,x,&x)
	//	fmt.Printf("%c",x)
	//}
	fmt.Println(cwd)





}
