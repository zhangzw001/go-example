package main

import (
	"log"
	"net"
	"strconv"
)

type ServerOpts struct {
	Host 	string
	Port 	int
}

func NewServer(s *ServerOpts ) {
	opts := ServerOpts{s.Host,s.Port}
	l := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	listener , err := net.Listen("tcp",l)
	if err != nil { log.Fatal( err )}
	defer listener.Close()


}

func main() {

}

