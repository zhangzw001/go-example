package main

import "fmt"

func main() {
	//ch := make(chan int , 3)
	//go func() {
	//	ch <- 1
	//	ch <- 2
	//}()
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(cap(ch))

	fmt.Println(mirroredQuery())
}


func mirroredQuery() []string {
	//如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。和垃圾变量不同，泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。
	responses := make(chan string , 3 )
	var l []string
	go func() { responses <- request("asia.gopl.io")}()
	go func() { responses <- request("eurepo.gopl.io")}()
	go func() { responses <- request("americas.gopl.io")}()
	// 方法1
	//for x := 0 ; x < 3 ; x ++ {
	//	l = append(l, <-responses)
	//}
	// 方法2
	for x := range responses {
		l = append(l, x)
	}
	close(responses)
	return l
}

func request(hostname string ) ( response string ) {
	return hostname
}