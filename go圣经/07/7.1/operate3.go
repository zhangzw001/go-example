package main

import (
	"fmt"
)

type Oper interface {
	Operate() int
}


type object struct {
	m int
	n int
}

type add struct {
	object
}

func (p *add) Operate() int {
	return p.m+p.n
}

type sub struct {
	object
}

func (s *sub) Operate() int {
	return s.m-s.n
}

func allOperate(o Oper) int {
	return o.Operate()
}

func mathOperate(s string ,a,b int ) int {
	switch s {
	case "+":
		add := &add{object{a,b}}
		return allOperate(add)
	case "-":
		sub := &sub{object{a,b}}
		return allOperate(sub)
	default:
		fmt.Println("不支持的运算")
	}
	return 0
}

func main() {
	//多态, 调用同一个方法, 实现不同的表现
	fmt.Println(mathOperate("-",1,2))
	fmt.Println(mathOperate("+",1,2))


}


