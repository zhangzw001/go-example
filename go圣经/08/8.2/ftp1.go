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

const uploadDir = "/tmp/ftp1/"

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

func conn(conn net.Conn) {
	defer conn.Close()
	for {
		io.WriteString(conn, "> ")
		r := bufio.NewReader(conn)
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
			//io.Copy(os.Stdout,c)
			msg, _ := ioutil.ReadAll(stdout)
			conn.Write(msg)
		case "upload":
			Upload(cmdList[1:],conn)
		case "close":
			conn.Close()
		case "help":
			io.WriteString(conn, "help\t帮助\n")
			io.WriteString(conn, "ls /\t列出目录内文件\n")
			io.WriteString(conn, "close\t关闭连接\n")
		default:
			io.WriteString(conn, "不支持的命令\n")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
// 支持多文件并发上传
func Upload(files []string, conn net.Conn) {
	//defer conn.Close()
	for _ , file := range files {
		go upload(file,conn)
	}
}
func upload(file string , conn net.Conn) {
		lfile := strings.Split(file,"/")
		cfile := strings.Join([]string{uploadDir,lfile[len(lfile)-1]},"")
		f , err := os.Create(cfile)
		if err != nil {
			log.Fatalf("文件创建失败: %v\n",cfile)
		}


		// 正常来说这里应该是写个客户端读取文件, 我这里仅测试, 写成本机读取 server(本机) 存储了
		rFile , _ := os.Open(file)
		readFile := bufio.NewScanner(rFile)
		for readFile.Scan() {
			f.WriteString(readFile.Text())
			f.WriteString("\n")
		}
		io.WriteString(conn,"文件上传成功\n")
		defer f.Close()
}

func main() {
	opts := ServerOpts{"localhost", 8000}
	NewServer(&opts)
}
