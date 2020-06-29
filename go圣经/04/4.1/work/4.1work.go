package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//fmt.Println(compareSha256("x","x"))
	fmt.Println(compareSha256("x","X"))
	fmt.Println(compareSha256("a","b"))
}



func compareSha256(str1 string,str2 string)int{
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	//循环字节数组
	fmt.Println(len(a))
	for i:=0;i<len(a);i++{
		//1个字节8个bit,移位运算，获取每个bit
		for m:=1;m<=8;m++{
			//比较每个bit是否相同
			if (a[i] >> uint(m))!=(b[i] >>uint(m)){
				num++
			}else {
				//fmt.Printf("%d,%d",i,m)
			}
		}
	}
	return num
}