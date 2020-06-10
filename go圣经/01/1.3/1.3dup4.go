package main

import (
	"bufio"
	"fmt"
	"os"
)

//type Pair struct {
//	Key   string
//	Value int
//}
//
//type PairList []Pair

func main() {
	//var p1 PairList

	map1 := make(map[string]map[string]int)

	files := []string{"./go圣经/01/1.3/1.3bufio-learn1.txt","./go圣经/01/1.3/1.3bufio-learn2.txt"}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil { fmt.Println(err)}


		scanner := bufio.NewScanner(f)
		map1[arg] = make(map[string]int)
		for scanner.Scan() {
			map1[arg][scanner.Text()]++
		}
	}

	//fmt.Println(map1)
	for k1,_ := range map1 {
		fmt.Println("文件名: ",k1)
		for k,v := range map1[k1] {
			fmt.Printf("%d\t\t%s\n",v,k)
		}
	}

}

