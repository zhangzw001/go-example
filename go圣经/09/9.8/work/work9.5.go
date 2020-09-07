package main

import (
	"fmt"
	"runtime"
	t "time"
)

//练习 9.5: 写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。
//这样的程序每秒可以支持多少次通信？

func echo() {
	ping := make(chan string)
	pong := make(chan string)

	m,n := 0,0

	timeout := make(chan struct{})
	go func() {
		t.Sleep(1 * t.Second)
		timeout <- struct {}{}
	}()
	go func() {
		for {
			ping <- "ping"
			pong <- "pong"
		}
	}()
	loop:
		for {
			select {
			case p := <-ping:
				m++
				fmt.Println(p)
			case q := <-pong:
				n++
				fmt.Println(q)
			case <-timeout:
				fmt.Println(m,n)
				fmt.Println("timeout")
				break loop
			default:
			}
		}

}
func main() {
	runtime.GOMAXPROCS(1)
	echo()
}
