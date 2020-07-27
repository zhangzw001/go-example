package main

import (
	"fmt"
	"math"
)

type Point struct{X,Y float64}

type Path []Point

func Distance(p , q Point ) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func (p Point ) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Sum() float64{
	return p.X+p.Y
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
func main() {
	p1 := Point{2,3}
	p2 := Point{4,6}
	//1 可以看到，上面的两个函数调用都是Distance，但是却没有发生冲突。
	//第一个Distance的调用实际上用的是包级别的函数geometry.Distance，
	//而第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法。
	fmt.Println(Distance(p1,p2))
	fmt.Println(p1.Distance(p2))
	fmt.Println(p1.Sum())

	//2 Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方法。在能够给任意类型定义方法这一点上，Go和很多其它的面向对象的语言不太一样。因此在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型(译注：这个例子里，底层类型是指[]Point这个slice，Path就是命名类型)不是指针或者interface。
	p3 := Point{6 ,9 }
	path := Path{p1,p2,p3,p1}
	path.Distance()
	//让我们来调用一个新方法，计算三角形的周长：
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}