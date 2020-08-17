package main

import (
	"fmt"
	"time"
)

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
	j := f()
	for i := 0 ; i < n ; i ++ {
		j = f()
	}
	return j
}

// 很慢的递归写法
func Fib2(n int ) int {
	if n < 2 { return 1 }
	return Fib2(n-1)+Fib2(n-2)

}

//
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}
func main() {
	go spinner(100 * time.Millisecond)
	n := 40
	fmt.Println(Fib(n))
	fmt.Printf("\rFibonacci(%d) = %d\n",n,Fib2(n))
}