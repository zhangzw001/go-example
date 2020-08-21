package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("获取到abort信号...")
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(<-tick)
		select {
		case <-tick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("aborted...")
			return
		default:

		}
	}

	//launch()
}
