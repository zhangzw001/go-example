package main

import "fmt"

var global *int

func main() {

	e()
	g()
	fmt.Println(*global)
}

func e() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
