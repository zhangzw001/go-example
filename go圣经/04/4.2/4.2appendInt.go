package main

import "fmt"

func appendInt(s []int , i ...int )  []int {

	var rst []int
	zlen := len(s)+len(i)

	if zlen <= cap(s) {
		rst = s[:zlen]
	}else {
		//zcap := zlen
		//if zcap < 2*len(s) {
		//	zcap = 2 * len(s)
		//}
		zcap := zlen
		if zcap > 2{
			zcap = 2 * len(s)
		}
		rst = make([]int ,zlen,zcap)
		copy(rst,s)
	}
	//rst[zlen-1] = i
	//fmt.Printf("%T\n",i)
	copy(rst[len(s):],i)
	return rst
}

//删除保持原来顺序
func delIndex(s []int , i int ) []int {
	copy(s[i:],s[i+1:])
	return s[:len(s)-1]
}
//如果删除元素后不用保持原来顺序的话，我们可以简单的用最后一个元素覆盖被删除的元素
func delIndex2(s []int , i int ) []int {
	n := len(s)
	s[i] = s[n-1]
	return s[:n-1]
}
func main() {
	//s1 := make([]int , 3, 10)
	//s3 := []int{1,2,3}
	//copy(s1, s3)		//copy 不会新建新的内存空间，由它原来的切片长度决定
	////s1 = s3
	//fmt.Println(s1)
	//fmt.Println(s3)
	//s2 := appendInt(s1,4)
	//fmt.Println(s2)
	//fmt.Println(cap(s1),len(s2),cap(s2))
	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%2d\t%v\n", i, cap(y), y)
	//	x = y
	//}
	//s4 := []int{1,2,3}
	//s5 := append(s4,5,6)
	//s6 := appendInt(s4,5,6)
	//fmt.Println(s5,s6)
	//
	//copy(s4[2:],[]int{8})
	//fmt.Println(s4)


	// delIndex 测试
	s7 := []int{1,2,3,4,5,6,7}
	s7 = delIndex(s7,2)
	fmt.Println(s7)
	// delIndex2 测试
	s8 := []int{1,2,3,4,5,6,7}
	s8 = delIndex2(s8,4)
	fmt.Println(s8)
}