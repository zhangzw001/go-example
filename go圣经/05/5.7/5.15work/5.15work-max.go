package main

import (
	"fmt"
	"log"
)
//练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
func main() {
	fmt.Println(Sort(9,4,2,3,5,99,1333))
	fmt.Println(Max(9,4,2,3,5,99,1333))
	fmt.Println(Min(9,4,2,3,5,99,1333))
}
func Sort(vals ...int ) (int ,int ){
	if len(vals) == 0  {
		log.Fatal("is nil")
	}
	n := len(vals)
	for i :=0 ; i < n ; i ++ {

		for j := i ; j < n ; j ++ {
			if vals[i] < vals[j] {
				vals[i],vals[j] = vals[j],vals[i]
			}
		}
	}
	return vals[0],vals[n-1]
}

func Max(vals ...int) int {
	if len(vals) == 0  {
		log.Fatal("is nil")
	}
	var max int
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}
func Min(vals ...int) int {
	if len(vals) == 0  {
		log.Fatal("is nil")
	}
	min := vals[0]
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min
}
