package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

//练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	ch := make(chan string )
	input := bufio.NewScanner(c)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Done()
		for {
			select {
			case <-time.After(3 * time.Second):
				c.Close()
			case input := <-ch:
				echo(c, input, 1*time.Second)
			}
		}
	}()
		for input.Scan() {
			ch <- input.Text()
		}
	wg.Wait()

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		//go handleConn(conn)
		go handleConn(conn)
	}
}



//网上别人写的
func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	var message = make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			//case <-time.After(6 * time.Second):
			case <-time.After(6 * time.Second):
				c.Close()
			case mes := <-message:
				wg.Add(1)
				echo(c,mes,1*time.Second)
				//go func(c net.Conn, shout string, delay time.Duration) {
				//	defer wg.Done()
				//	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
				//	time.Sleep(delay)
				//	fmt.Fprintln(c, "\t", shout)
				//	time.Sleep(delay)
				//	fmt.Fprintln(c, "\t", strings.ToLower(shout))
				//	//ch<-struct{}{}
				//}(c, mes, 1*time.Second)

			}
		}
	}()
	for input.Scan() {
		text := input.Text()
		message <- text
	}

	wg.Wait()
	//cw := c.(*net.TCPConn)
	//cw.CloseWrite()

	c.Close()
}