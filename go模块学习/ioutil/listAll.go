package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func listAll(path string, n int ) {
	r, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print(err)
	}
	for _, info := range r {
		if info.IsDir() {
			for i := n ; i > 0 ; i-- {
				fmt.Printf("|\t")
			}
			//fmt.Println(info.Name(),"/")
			fmt.Printf("%v/\n",info.Name())
			listAll(path + "/" + info.Name(), n+1)
		}else {
			for i := n ; i > 0 ; i-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
			//fmt.Println(info.Name(),"/")
		}
	}
}


func main() {
	//var dir []string
	//if len(os.Args) == 1 {
	//	dir =
	//}else {
	//	dir := os.Args[1]
	//}
	//listAll(dir,0)


	//dir := os.Args[1]
	dir := "/Users/zhangzw/lrzsz/logs/sql"
	listAll(dir,0)
}
