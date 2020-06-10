package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {

	//scanner := bufio.NewScanner(strings.NewReader("姓名;年龄;籍贯;性别\r\nTom,15,China,male"),)
	//scanner.Scan()
	//fmt.Println("第一次读取:",scanner.Text())
	////1
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}
	//
	////2
	//scanner.Split(bufio.ScanWords)
	//scanner.Split(bufio.ScanLines)
	//scanner.Split(ScanItems)
	//for scanner.Scan(){
	//	//按照默认的行读取,每次读取一行
	//	fmt.Println("循环内读取:",scanner.Text())
	//	fmt.Println("---")
	//}
	start := time.Now()
	counts := make(map[string]int)
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	file,err :=os.Open("./go圣经/01/1.3/1.3bufio-learn1.txt")

	if err != nil {
		fmt.Println("open file err: ",err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	for k,v := range counts {
		fmt.Printf("%d\t%s\n",v,k)
	}
	fmt.Println(len(counts))
	fmt.Printf("%.6f\n",time.Since(start).Seconds())

}

func ScanItems(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexAny(data, ",;"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
