package main

import "fmt"


func main() {
	//var p = f()
	fmt.Println(f() == f())

	//
	v :=1
	incr(&v)
	fmt.Println(incr(&v))
}

//
func f() *int {
	v := 1
	fmt.Println(&v)
	return &v
}


//
func incr(p *int) int {
	*p ++
	return *p
}