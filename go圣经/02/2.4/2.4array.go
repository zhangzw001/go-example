package main

import "fmt"

type inter interface {
}
func main() {

	fmt.Println(gcd(63, 81))
	n := 10
	for i:=0 ; i<n ; i++ {
		fmt.Printf("%d ",fib(i))
	}
	fmt.Println()
	for i:=0 ; i<n ; i++ {
		fmt.Printf("%d ",fib2(i))
	}
	fmt.Println()

	var  a inter
	a = 1
	if v,ok := a.(int) ; ok {
		fmt.Println(v)
	}

}


func gcd(x , y int ) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int ) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int ) int {
	x, y := 0, 1
	for i:=0 ;i < n ; i++ {
		x, y = y, x+y
	}
	return x
}