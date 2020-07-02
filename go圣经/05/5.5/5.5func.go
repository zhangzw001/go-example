package main

import (
	"fmt"
	"strings"
)

func main() {
	//函数类型的零值是nil。调用值为nil的函数值会引起panic错误：
	//函数值可以与nil比较：
	//但是函数值之间是不可比较的，也不能用函数值作为map的key。
	var f func(int) int
	if f != nil {
		f(1)
	}

	//函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。标准库中包含许多这样的例子。下面的代码展示了如何使用这个技巧。strings.Map对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}
func add1(r rune) rune {return r + 1 }
