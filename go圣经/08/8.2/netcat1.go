package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp","localhost:8000")
	defer conn.Close()
	mustCopy(os.Stdout,conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src ); err != nil {
		log.Fatal(err )
	}
}
