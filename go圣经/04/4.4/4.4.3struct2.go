package main

import "fmt"

type Point struct {
	X, Y int
}
type point2 struct {
	X, Y int
}

//type Circle struct {
//	Center Point
//	Radius int
//}
//
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type circle struct {
	point2
	Radius int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type Wheel2 struct {
	circle
	Spokes int
}

func main() {
	////
	//var w Wheel
	//w.Circle.Center.X = 8
	//w.Circle.Center.Y = 8
	//w.Circle.Radius = 5
	//w.Spokes = 20

	////
	//var w Wheel
	//w.X = 8            // equivalent to w.Circle.Point.X = 8
	//w.Y = 8            // equivalent to w.Circle.Point.Y = 8
	//w.Radius = 5       // equivalent to w.Circle.Radius = 5
	//w.Spokes = 20

	// 结构体字面值初始化必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此是等价的：
	w := Wheel{Circle{Point{8,8},5},20}
	w2 := Wheel{
		Circle: Circle{
			Point: Point{X:8, Y:8},
			Radius: 5,
		},
		Spokes: 20,
	}
	//需要注意的是Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
	fmt.Printf("%#v \n%#2v\n",w,w2)
	fmt.Printf("%v \n%2v\n",w,w2)

	w.X = 42
	fmt.Printf("%#v\n", w)
	//因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
	//同时，因为成员的名字是由其类型隐式地决定的，所有匿名成员也有可见性的规则约束。
	//在上面的例子中，Point和Circle匿名成员都是导出的。
	//即使它们不导出（比如改成小写字母开头的point和circle），我们依然可以用简短形式访问匿名成员嵌套的成员

	w3 := Wheel2{
		circle: circle{
			point2: point2{
				X: 6,
				Y: 6,
			},
		},
		Spokes: 10,
	}
	fmt.Printf("%#v\n",w3)
	fmt.Println(w3.X,w3.circle.point2.X)
	//但是在包外部，因为circle和point没有导出不能访问它们的成员，因此简短的匿名成员访问语法也是禁止的。

}