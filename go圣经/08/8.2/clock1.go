package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener , err := net.Listen("tcp", "localhost:8000")
	if err != nil { log.Fatal(err )}
	for {
		conn, _ := listener.Accept()
		//conn.Read()
		handleConn(conn)
	}
}


func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05 UTC-0700\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

