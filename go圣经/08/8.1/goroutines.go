package main

import "fmt"

// 快的写法, 但是空间 换时间
func fib() func() int{
	var x,y int
	return func() int {
		if x < 2 { x = 1 }
		x , y = x+y, x
		return y
	}
}

func Fib(n int ) int {
	f := fib()
	var i,j int
	for i = 0 ; i < n ; i ++ {
		j = f()
	}
	return j
}

// 很慢的递归写法
func Fib2(n int ) int {
	if n < 2 { return 1 }
	return Fib2(n-1)+Fib2(n-2)

}
func main() {
	fmt.Println(Fib(1))
	fmt.Println(Fib2(40))
}