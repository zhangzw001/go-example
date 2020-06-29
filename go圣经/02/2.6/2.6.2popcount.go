package main

import (
	"fmt"
	"github.com/zhangzw001/go-example/go圣经/02/2.6/popcount"
)

func main(){
	//fmt.Println(popcount.PopCount(10000000))
	//fmt.Printf("%b\n",10000000)
	fmt.Println(popcount.PopCount(7))
	fmt.Println(popcount.PopCount(512))
	fmt.Println(popcount.PopCountByLoop(7))
	//fmt.Println(popcount.PopCountByLoopShift(10000000))
	//fmt.Println(popcount.PopCountByAnd(10000000))
	//fmt.Println(popcount.PopCountByAnd(99))
	//fmt.Printf("%b\n",99)


	//// 8 	-> 1000&0111 -> 0000
	//// 16 	-> 10000&01111 -> 00000
	//fmt.Printf("%b",math1(8))
}


func math1(x int) bool {
	//如果一个数是2的n次方，那么这个数用二进制表示时其最高位为1，其余位为0。
	if x&(x-1) == 0 {
		return true
	}else {
		return false
	}
}
