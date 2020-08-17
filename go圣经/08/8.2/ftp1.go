package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
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
		io.WriteString(c, "> ")
		r := bufio.NewReader(c)
		list, _, _ := r.ReadLine()
		cmd := strings.Join([]string{string(list)}, "")
		//io.WriteString(c,strings.Join([]string{string(list)},""))
		cmdList := strings.Split(cmd, " ")
		switch cmdList[0] {
		case "ls":
			cmd := exec.Command(cmdList[0], cmdList[1:]...)

			stdout, err := cmd.StdoutPipe()
			if  err != nil {
				log.Fatal(err)
			}
			defer stdout.Close()

			err = cmd.Start()
			//err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			msg, _ := ioutil.ReadAll(stdout)
			c.Write(msg)
		case "close":
			c.Close()
		case "help":
			io.WriteString(c, "help\t帮助\n")
			io.WriteString(c, "ls /\t列出目录内文件\n")
			io.WriteString(c, "close\t关闭连接\n")
		default:
			io.WriteString(c, "不支持的命令\n")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	opts := ServerOpts{"localhost", 8000}
	NewServer(&opts)
}
