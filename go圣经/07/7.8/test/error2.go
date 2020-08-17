package test

import "fmt"

type p struct {
	//s string
}


func error2() bool {
	p1 := &p{}
	p2 := &p{}
	fmt.Println(p1,p2)
	if p1 == p2 {
		return true
	}
	return false
}