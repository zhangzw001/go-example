package main

import (
	"fmt"
	github "github.com/zhangzw001/go-example/go圣经/04/4.5/json-github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	////
	//修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
	//当前时间
	now := time.Now().Unix()
	//前一个月
	preMonth := now - 30*24*3600
	// 前一年
	preYear := now - 365*24*3600

	var notMonth []*github.Issue
	var notYear []*github.Issue
	var overYear []*github.Issue

	for _, item := range result.Items {
		createTime := item.CreatedAt.Unix()
		if createTime > preMonth {
			notMonth = append(notMonth,item)
			continue
		}
		if createTime < preMonth && createTime > preYear {
			notYear = append(notYear, item)
			continue
		}
		overYear = append(overYear, item)
		//fmt.Println("issues(不到一个月):")
		for _, item := range notMonth {
			fmt.Printf("不到一个月:  #%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}

		//fmt.Println("issues(不到一年):")
		//for _, item := range notYear {
		//	fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		//}
		//fmt.Println("issues(超过一年):")
		//for _, item := range overYear {
		//	fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		//}
	}
	/////
	//
	//
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s %s\n",
	//		item.Number, item.User.Login, item.Title, item.CreatedAt)
	//}
}