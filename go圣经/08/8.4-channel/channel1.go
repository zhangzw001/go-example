package main

import "fmt"

func writeC(c chan int )  {
	c <- 1
}
func main() {
	ch := make(chan int)
	defer close(ch)
	// 写到通道
	// ch <- 1 //  all goroutines are asleep - deadlock!, 因为写入必须等 读了才会 释放阻塞, 但是main goroutine 已经被阻塞,无法继续执行 读( x :=<- ch)操作, 导致一直阻塞中
	go writeC(ch)	// 开启一个新的goroutine 来写入
	x :=<- ch
	fmt.Println(x)
	//ch := make(chan int, 10)
	//ch <- 11
	//ch <- 12
	//
	//close(ch)
	//
	//for x := range ch {
	//	fmt.Println(x)
	//}
	//
	//x, ok := <- ch
	//fmt.Println(x, ok)
}
