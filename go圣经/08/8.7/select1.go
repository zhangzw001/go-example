package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		timeout <- struct{}{}
	}()

	ch := time.Tick(1 * time.Second)
	for {
		select {
		case <-ch:
			fmt.Println(<-ch)
			fmt.Println(time.Now())
			time.Sleep(8 * time.Second)
			fmt.Println(time.Now())
		case <-timeout:
			fmt.Println("timeout!")
			return
		}
	}
}
