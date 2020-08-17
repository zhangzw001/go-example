package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c,"\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))

}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c,input.Text(),1*time.Second)
		//echo(c,input.Text(),1*time.Second)
	}
}

func main() {
	listener , _ := net.Listen("tcp","localhost:8000")
	for {
		conn, _ := listener.Accept()
		handleConn(conn)
	}
}