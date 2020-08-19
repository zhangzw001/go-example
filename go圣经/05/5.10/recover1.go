package main

import "fmt"

func A(x int ) {
	fmt.Println("A")
}

func B(x int ) {
	defer func() {
		//recover()
		// 打印panic, 只有在有错误的时候打印
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var a [2]int
	a[x] = 1		//x 大于1 都会导致panic
}


func C(x int ) {
	fmt.Println("C")
}

func main() {
	A(1)
	B(1)
	C(2)
}