package main

import (
	"fmt"
	"time"
)

type point struct {
	X,Y int
}
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type address struct {
	hostname 	string
	port 		int
}
func main() {
	// 第一种字面值, 顺序一一对应
	p := point{1,2}
	fmt.Println(p)
	// 第二种字面值初始化, 针对成员进行初始化, 忽略的成员则是默认0值
	p2 := point{Y: 1}
	fmt.Println(p2)

	//
	fmt.Println(Scale(point{1,2},10))

	//
	employee := Employee{Name:"张三",Salary:10000}
	fmt.Println(Bonue(&employee, 3))
	AwardAnnualRaise(&employee, 3)
	fmt.Println(employee.Salary)

	//因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：
	pp := &point{1,2}
	fmt.Println(pp)

	// 结构体的比较
	hits := make(map[address]int)
	hits[address{"golang.org",443}]++


}


func Scale(p point, factor int ) point {
	return point{p.X * factor , p.Y * factor}
}

// 如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回，
func Bonue(e *Employee, percent int ) int {
	return e.Salary * percent /100
}
//如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
func AwardAnnualRaise(e *Employee, percent int ) {
	e.Salary = e.Salary * (100 + percent) / 100
}