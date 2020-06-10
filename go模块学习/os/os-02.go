package main


import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./os-02.txt")
	if err != nil {
		fmt.Println("open file err: ", err.Error())
		return
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		//str, err := r.ReadString('\n')
		data, err := r.ReadBytes('\n')
		// 读取文件末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err :" ,err.Error())
		}
		for i,_ := range data {
			fmt.Println(data[i])
			break
		}
		break
		//fmt.Printf("%v",string(data))

	}
}