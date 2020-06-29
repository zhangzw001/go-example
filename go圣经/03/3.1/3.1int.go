package main

import "fmt"

func main(){
	//&      位运算 AND
	//|      位运算 OR
	//^      位运算 XOR
	//&^     位清空 (AND NOT)
	//<<     左移
	//>>     右移
	//var u uint8 = 255
	//fmt.Println(u, u+1, u*u) // "255 0 1"
	//
	//var i int8 = 127
	//fmt.Println(i, i+1, i*i) // "127 -128 1"
	//
	//var j uint16 = 65535
	//fmt.Println(j, j+1, j*j) // "65535 0 1"
	//
	////
	//var x uint8 = 1<<1 | 1<<5
	//var y uint8 = 1<<1 | 1<<2
	//
	//fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	//fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}
	//
	//fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	//fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	//fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	//fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	//for i := uint(0); i < 10; i++ {
	//	fmt.Println(x&(1<<i))
	//	if x&(1<<i) != 0 { // membership test
	//		fmt.Println(i) // "1", "5"
	//	}
	//}
	//fmt.Println(1<<8,1024>>8)
	////
	//medals := []string{"gold", "silver", "bronze"}
	//for i := 0 ; i<len(medals) ; i++ {
	//	fmt.Println(i)
	//}
	//for i := len(medals) - 1 ; i>=0; i-- {
	//	fmt.Println(medals[i])
	//}
	f := 1e100  // a float64
	i := int(f) // 结果依赖于具体实现
	fmt.Println(i)
	//请注意fmt的两个使用技巧。通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。
	//第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	o := 0666
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]o %#[1]o , %d %[2]x %#[2]x %#[2]X\n", o,x) // "438 666 0666 , 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF"

	//单个字符
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	s1 := "test"
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q %% %q\n", newline,s1)       // "10 '\n'"
}
