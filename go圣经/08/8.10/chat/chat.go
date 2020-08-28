package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client chan<- string

var user = make(map[string]string)
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func Broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}

}


func HandleConn(conn net.Conn) {
	ch := make(chan string ) //outgoing client messages
	go clientWriter(conn, ch)


	who := conn.RemoteAddr().String()
	if user[who] != "" {
		messages <- user[who] + " welcome back"
	}else {
		fmt.Println("等待输入昵称")
		buf := bufio.NewReader(conn)
		tmp,_ := buf.ReadString('\n')
		user[who] = strings.ReplaceAll(tmp,"\n","")
		who = user[who]
	}
	fmt.Println(who,user)
	ch <- "You are " + who
	messages <- who + "  has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + "  has left"
	conn.Close()
}

func clientWriter(conn net.Conn , ch <- chan string ) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}