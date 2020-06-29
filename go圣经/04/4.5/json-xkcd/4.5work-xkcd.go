package main

//练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Xkcd struct {
	Num 	int 	`json:"num"`
	Img 	string  `json:"img"`
	Title 	string  `json:"title"`
	Transcript	string `json:"transcript"`
}

func FetchChan(url string , ch chan <- string) {
	var result Xkcd
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	titles := strings.Split(result.Transcript, " ")
	for _, v := range titles {
		if v == string(os.Args[1]) {
			ch <- result.Img
		}
	}

}

func Fetch(url string) Xkcd {
	var result Xkcd
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	titles := strings.Split(result.Transcript, " ")
	for _, v := range titles {
		if v == string(os.Args[1]) {
			return result
			break
		}
	}
	return Xkcd{}

}

func main() {
		var urls []string
		for i:=1 ;i < 50 ; i++ {
			url := fmt.Sprintf("https://xkcd.com/%d/info.0.json",i)
			urls = append(urls,url)
		}

		// 网上channel方式(但是每次只返回一条)
		ch := make(chan string )
		for _, url := range urls {
			go FetchChan(url , ch )
		}
		fmt.Println(<-ch)

		// 自己写的返回内容添加到map中, 可以返回多个
		xkcdlist := make(map[int]string)
		for _, url := range urls {
				tmp := Fetch(url)
				//fmt.Println(tmp.Img,tmp.Num)
				xkcdlist[tmp.Num] = tmp.Img
		}
		for k,_ := range xkcdlist {
			if k == 0 { continue }
			fmt.Printf("https://xkcd.com/%d/info.0.json : %v\n",k,xkcdlist[k])

		}
}