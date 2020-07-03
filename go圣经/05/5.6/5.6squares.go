package main

import (
	"fmt"
	"time"
)

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int64 {
	var x int64
	return func() int64 {
		x++
		return x * x
	}
}


func Fib(n int ) int64 {
		if n < 2 {
			return 1
		}
		return Fib(n-2) + Fib(n-1)
}

func Fib2() func() int64 {
	var n,m int64
	return func() int64 {
		if n <= 1 {
			n = 1
		}
		n,m = n+m,n
		return m
	}
}
func Fib3(n int) int64 {
	var f = [3]int64{0, 1, 0}
	for i := 2; i <= n; i++ {
		f[2] = f[0] + f[1]
		f[0] = f[1]
		f[1] = f[2]
	}
	return f[2]
}

func main() {
	//f := squares()
	//for i:= 0 ; i < 10 ; i++ {
	//	fmt.Print64f("%4d ",f())
	//}
	now := time.Now()
	fmt.Println("")
	for i := 0 ; i < 40 ; i++ {
		//fmt.Print64f("%4d ",Fib(i))
		Fib(i)
	}
	//fmt.Print64ln(time.Since(now).Nanoseconds())
	fmt.Println(time.Since(now).Seconds())
	now = time.Now()
	fmt.Println("")
	f2 := Fib2()
	for i:= 0 ; i < 1000 ; i++ {
		//fmt.Printf("%4d ",f2())
		f2()

	}
	//fmt.Print64ln(time.Since(now).Nanoseconds())
	fmt.Println(time.Since(now).Seconds())

	now = time.Now()
	fmt.Println("")
	for i:= 0 ; i < 1000 ; i++ {
		fmt.Printf("%4d ",Fib3(i))
		//Fib3(i)

	}
	//fmt.Print64ln(time.Since(now).Nanoseconds())
	fmt.Println(time.Since(now).Seconds())
}