package main

import (
	"fmt"
)

type Operate2 struct {}

func (p Operate2) Math(a, b float64, op string ) float64 {
	switch op {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		fmt.Println("仅支持输入 + - * /")
	}
	return 0
}


func main() {
	var p Operate2

	for _,i := range `+-*/` {
		fmt.Println(p.Math(1,2,string(i) ))
	}
}