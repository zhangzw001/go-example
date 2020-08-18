package main

import "fmt"

// out 只能发送, 不能接收
func counter(out chan<- int) {
	for x := 0 ; x < 10 ; x ++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int , in <-chan int ) {
	for x := range in  {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int ) {
	for x := range in {
		fmt.Println(x )
	}
}

func main(){
	//调用counter（naturals）时，naturals的类型将隐式地从chan int转换成chan<- int。调用printer(squares)也会导致相似的隐式转换，这一次是转换为<-chan int类型只接收型的channel。任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。这里并没有反向转换的语法：也就是不能将一个类似chan<- int类型的单向型的channel转换为chan int类型的双向型的channel。
	naturals := make(chan int )
	squares  := make(chan int )
	go counter(naturals)
	go squarer(squares,naturals)
	printer(squares)
}