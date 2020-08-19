package main

import "fmt"

type Oper interface {
	Operate(a,b int) int
}

type object struct {
	m int
	n int
}

type add struct {
	object
}

func (p *add) Operate(a , b int ) int {
	p.m , p.n = a, b
	return p.m+p.n
}

type sub struct {
	object
}

func (s *sub) Operate(a , b int) int {
	s.m, s.n = a, b
	return s.m-s.n
}


func main() {
	var o [2]Oper

	add1 := &add{}
	o[0] = add1
	fmt.Println(o[0].Operate(1,2))
	sub1 := &sub{}
	o[1] = sub1
	fmt.Println(o[1].Operate(1,2))


}
