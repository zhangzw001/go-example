package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//seed := time.Now().UTC().UnixNano()
	//fmt.Printf("Random seed: %d",seed)
	//r := rand.New(rand.NewSource(seed))
	//fmt.Println(0x1000)
	//fmt.Println(0x0FFF)
	//fmt.Println(r.Intn(0x0999-0x1000)+0x999)
	//fmt.Println(word.IsPalindrome2("été"))
	//storageMail.CheckQuota("zhangzw@boqii.com")
	//storageMail.GomailLearn1()
	//1.连接RabbitMQ,RabbitMQ的连接已经为我们抽象了socket的连接，同时为我们处理了协议版本号和身份认证等等
	conn, err := amqp.Dial("amqp://guest:guest@172.16.76.133:8081")
	if err != nil { log.Fatal(err)}
	defer conn.Close()
	//2.创建通道
	ch , err := conn.Channel()
	if err != nil { log.Fatal(err)}
	defer ch.Close()
	//3.声明队列并发送数据
	//在开始发送消息之前我们首先应该声明一个队列。声明队列之后我们就可以将消息发送至队列当中
	q, err := ch.QueueDeclare(
		"hello", //name
		false, //
		false, //
		false, //
		false, //
		nil,
		)
	if err != nil { log.Fatal(err)}
	body := "hello world!"

	for i :=0 ; i < 10 ; i ++ {
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Fatal(err)
		}
	}
}

