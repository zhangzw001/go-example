package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/hpcloud/tail"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)



func main(){

	redisHost := "172.16.76.194:30123"
	redisDb := 2


	c, err := redis.Dial("tcp",redisHost)
	if err != nil {
		fmt.Println("connect to redis error",err)
		return
	}
	c.Do("select",redisDb)
	defer c.Close()


	fileName := "./os-01.txt"
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
		string1 string
		count int
	)


	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		s := line.Text
		index := strings.Index(s,"the content")
		lens := len(s)
		if index == -1 {
			fmt.Println("没找到 the content")
		}else {
			//fmt.Println(index)
			//fmt.Println(s[index+12:])
			json1 := s[index+13:lens-1]
			if ! gjson.Valid(json1) {
				fmt.Println("invalid json")
				return
			}

			myudid := gjson.Get(json1,`udid`)

			string1 += " "+myudid.String()
			count ++

			fmt.Println(string1)

			if count > 10 {
				_, err = c.Do("pfadd", "udid", string1 )
				string1 = ""
				count = 0
				if err != nil {
					fmt.Println("redis pfadd failded:", err)
				}
			}
		}


	}
}

