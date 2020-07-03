package main

import "fmt"

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}


func fib(n int ) int {
		if n < 2 {
			return 1
		}
		return fib(n-2) + fib(n-1)
}

func fib2() func() int {
	var n int
	return func() int {
		n ++
		if n < 2 {return n}
		n = 2 * n  - 3
		return n
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	//for i:=0 ; i < 10 ; i++ {
	//	fmt.Println(fib(i))
	//}
	f2 := fib2()
	for i:= 0 ; i < 10 ; i++ {
		fmt.Println(f2())
	}
}