package main

import (
	"fmt"
	"reflect"
)


func main(){
	k:=40
	fmt.Println(k)
	fmt.Println(say("hello,world","lf"))
	fmt.Println(reflect.TypeOf(k))  //检查变量类型
	var pk *int
	pk = &k
	fmt.Printf("%p\n",pk)
	fmt.Printf("%p\n",&k)

	fmt.Println(&k)  // 获取变量在计算机内存中的地址，可在变量名前面加上&字符。
	// &k 引用的是变量k的值，值所在的内存地址
	showMemoryAddress(&k)  //返回的地址是相同的
}

func showMemoryAddress(x *int){     // *int参数类型位指向整数的指针，而不是整数
	fmt.Println(*x)  //本身就是指针，打印地址不需要  &这个符号，如果想使用指针指向的变量的值，而不是其内存地址，可在指针变量前面加上星号
	return
}

func say(param,tt string) string {
	return param + "--" + tt

}