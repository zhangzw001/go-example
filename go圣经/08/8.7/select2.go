package main

import "fmt"

func main() {

	ch := make(chan int,1)
	go func() {
		
		fmt.Println(<-ch)
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		//select {
		//case x := <-ch:
			//fmt.Println(x)
		//case ch <- i:
		//}
	}
}
