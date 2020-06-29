package main

import (
	"fmt"
	"sort"
)

func main() {

	map1 := make(map[string]int)
	map1["age"] = 18
	fmt.Println(map1)

	age1 := map[string]int {
		"zhangsan": 19,
		"lisi": 28,
		"wangwu": 44,
		"zhaoliu": 2,
		"aaaa": 10,
	}
	delete(age1,"age")
	fmt.Println(age1)

	s := make([]string,0,len(age1))
	for i,_ := range age1 {
		s = append(s,i)
	}
	sort.Strings(s)
	for _,j := range s {
		fmt.Printf("key: %8s, value: %4d\n",j,age1[j])
	}


	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
	ages = make(map[string]int)
	fmt.Println(ages == nil)    // "false"
	fmt.Println(len(ages) == 0) // "true"
	ages["aaa"] = 1
	fmt.Println(ages)


}
