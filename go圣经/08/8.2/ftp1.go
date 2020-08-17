package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type ServerOpts struct {
	Host string
	Port int
}

func NewServer(s *ServerOpts) {
	opts := ServerOpts{s.Host, s.Port}
	l := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))
	listener, err := net.Listen("tcp", l)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		c, _ := listener.Accept()
		go conn(c)
	}

}

func conn(c net.Conn) {
	defer c.Close()
	for {
		io.WriteString(c, ">")
		r := bufio.NewReader(c)
		list, _, _ := r.ReadLine()
		cmd := strings.Join([]string{string(list)}, "")
		//io.WriteString(c,strings.Join([]string{string(list)},""))
		cmdList := strings.Split(cmd, " ")
		switch cmdList[0] {
		case "ls":
			io.WriteString(c, "正在执行命令\n")
			cmd := exec.Command(cmdList[0], cmdList[1:]...)
			cmd.Run()
			msg, _ := ioutil.ReadAll(os.Stdout)
			c.Write(msg)
		default:
			io.WriteString(c, "不支持的命令")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	opts := ServerOpts{"localhost", 8000}
	NewServer(&opts)
}
