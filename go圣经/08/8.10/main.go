package main

import (
	"github.com/zhangzw001/go-example/go圣经/08/8.10/chat"
	"log"
	"net"
)

func main() {

	listener, _ := net.Listen("tcp","localhost:8000")

	go chat.Broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go chat.HandleConn(conn)
	}

}
