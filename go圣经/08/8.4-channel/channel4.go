package main

import (
	"fmt"
)

func main() {
	//ch := make(chan int , 3)
	//go func() {
	//	ch <- 1
	//	ch <- 2
	//}()
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(cap(ch))
	res := make(chan string ,3)
	go mirroredQuery(res)
	readChannel(res)
	//fmt.Println(mirroredQuery())
}


func mirroredQuery(responses chan string ) {
	//如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。和垃圾变量不同，泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。
	//responses := make(chan string, 3)
	 responses <- request("asia.gopl.io")
	 responses <- request("eurepo.gopl.io")
	 responses <- request("americas.gopl.io")
	 //responses <- request("americas2.gopl.io")
	 defer close(responses)
}

func readChannel(ch chan string ) {
	//////////
	for x := 0 ; x < 3 ; x ++ {
		fmt.Println(<-ch)
	}
	//////////
	//for x := range ch {
	//	//l = append(l, x)
	//	fmt.Println(x)
	//}
}
func request(hostname string ) ( response string ) {
	return hostname
}