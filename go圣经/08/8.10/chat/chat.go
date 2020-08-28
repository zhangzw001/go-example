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
	buf := bufio.NewReader(conn)
	tmp , _ := buf.ReadString('\n')
	nickname := strings.Replace(tmp,"\n","",-1)

	if user[nickname] != "" {
		fmt.Printf("老用户: %s : %s ,%s\n",nickname,user[nickname],who)
		messages <- user[nickname] + " welcome back, " + who
	}
	user[nickname] = who
	whoString := nickname+"@"+user[nickname]
	fmt.Println(who,user)
	ch <- "You are " + whoString
	messages <- whoString + "  has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- whoString + ": " + input.Text()
	}

	leaving <- ch
	messages <- whoString + "  has left"
	conn.Close()
}

func clientWriter(conn net.Conn , ch <- chan string ) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}