package main

import (
	"fmt"
	"image/color"
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

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	p1 := Point{2,3}
	p2 := Point{4,6}
	//6.1.1 可以看到，上面的两个函数调用都是Distance，但是却没有发生冲突。
	//第一个Distance的调用实际上用的是包级别的函数geometry.Distance，
	//而第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法。
	fmt.Println(Distance(p1,p2))
	fmt.Println(p1.Distance(p2))
	fmt.Println(p1.Sum())


	//6.1.2 Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方法。在能够给任意类型定义方法这一点上，Go和很多其它的面向对象的语言不太一样。因此在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型(译注：这个例子里，底层类型是指[]Point这个slice，Path就是命名类型)不是指针或者interface。
	//让我们来调用一个新方法，计算三角形的周长：
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
	fmt.Println(PathDistance(perim))
	//译注： 在外部调用包时, 如果我们要用方法去计算perim的distance，还需要去写全geometry的包名，和其函数名，但是因为Path这个变量定义了一个可以直接用的Distance方法，
	//所以我们可以直接写perim.Distance()。相当于可以少打很多字，作者应该是这个意思。因为在Go里包外调用函数需要带上包名，还是挺麻烦的。

	//6.2.1
	r := &Point{1,2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	(&p).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"
	//6.2.2 如果接收器p是一个Point类型的变量，并且其方法需要一个Point指针作为接收器，我们可以用下面这种简短的写法：
	p.ScaleBy(2) // = (*p).ScaleBy(2)
	fmt.Println(p)
	//编译器会隐式地帮我们用&p去调用ScaleBy这个方法。这种简写方法只适用于“变量”，包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。
	//我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：
	// cannot call pointer method on Point literal
	// cannot take the address of Point literal
	//Point{1,2}.ScaleBy(2)


	//6.2.3 以下调用都是ok的
	//不论是接收器的实际参数和其接收器的形式参数相同，比如两者都是类型T或者都是类型*T：
	Point{1,2}.Distance(p1)
	pptr.ScaleBy(2)
	//或者接收器实参是类型T，但接收器形参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：
	p.ScaleBy(2)
	//或者接收器实参是类型*T，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：
	pptr.Distance(p1)

	//1. 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
	//2. 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的内部，第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。熟悉C或者C艹的人这里应该很快能明白。

	// 6.3.1
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p3 = ColoredPoint{Point{1, 1}, red}
	var q3 = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p3.Distance(q3.Point)) // "5"
	p3.ScaleBy(2)
	q3.ScaleBy(2)
	fmt.Println(p3.Distance(q3.Point)) // "10"

	// 一个ColoredPoint并不是一个Point，但他"has a"Point，并且它有从Point类里引入的Distance和ScaleBy方法。
	//如果你喜欢从实现的角度来考虑问题，内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法，和下面的形式是等价的：
	//func (p ColoredPoint) Distance(q Point) float64 {
	//	return p.Point.Distance(q)
	//}
	//
	//func (p *ColoredPoint) ScaleBy(factor float64) {
	//	p.Point.ScaleBy(factor)
	//}


	//6.4.1 将p4.Distance(q4) 拆分为两步
	p4 := Point{1,2}
	q4 := Point{4,6}

	distanceFromP := p4.Distance
	fmt.Println(distanceFromP(q4))

	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p4.ScaleBy
	scaleP(2)
	fmt.Println(p4)
	scaleP(3)
	fmt.Println(p4)
	scaleP(10)
	fmt.Println(p4)

	//当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，这种函数会将其第一个参数用作接收器，
	//所以可以用通常(译注：不写选择器)的方式来对其进行调用：
	// 译注：这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
	// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
	// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
	// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
	p41 := Point.Distance
	fmt.Println(p41(p4,q4))
	fmt.Println(p4.Distance(q4))

}