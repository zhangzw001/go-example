package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port",8000,"listen port")
	listener , err := net.Listen("tcp", "localhost:"+port)
	if err != nil { log.Fatal(err )}
	for {
		conn, _ := listener.Accept()
		//conn.Read()
		go handleConn2(conn)
	}
}


func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05 UTC-0700\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

