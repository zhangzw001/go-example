package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn , _ := net.Dial("tcp","localhost:8000")
	conn.Close()
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout,conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn,os.Stdin)
	//conn.Close()
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done

}


func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src ); err != nil {
		log.Fatal(err )
	}
}
