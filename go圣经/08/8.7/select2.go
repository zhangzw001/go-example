package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
			//select {
			//case x := <-ch:
			//fmt.Println(x)
			//case ch <- i:
			//}
		}
		defer close(ch)
	}()

	//for i := range ch {
	//	fmt.Println(i)
	//}
	go func() {
		for {
			a , ok := <-ch
			if !ok {
				return
			}
			fmt.Println(a)
		}
	}()
	time.Sleep(6 * time.Second)
}
