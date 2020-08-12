package main

import (
	"fib"
	"fmt"
)

func main() {
	for i :=0 ; i< 40; i++ {
		fmt.Printf("%d,",fib.Fib(i))
	}
	fmt.Println("")
	for i :=1 ; i< 90; i++ {
		fmt.Printf("%d,",fib.Fib2(i))
	}
	fmt.Println("")
	f3 := fib.Fib3()
	for i :=1 ; i< 90; i++ {
		fmt.Printf("%d,",f3())
	}

}
