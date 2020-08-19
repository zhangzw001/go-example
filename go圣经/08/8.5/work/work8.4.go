package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c,"\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))

}

func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		//go echo(c,input.Text(),1*time.Second,&wg)
		go echo(c,input.Text(),1*time.Second,&wg)
	}
	wg.Wait()
	defer c.Close()

}

func main() {
	listener , _ := net.Listen("tcp","localhost:8000")
	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}
