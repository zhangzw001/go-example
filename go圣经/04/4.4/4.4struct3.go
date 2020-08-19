package main

import "fmt"

type Person struct {
	Name 	string
	Id 		string
	Age 	string
}

type Student struct {
	Score	int
	*Person
}

type Teacher struct {
	Office	string
	Person
}

func main() {

	stu := new(Student)
	stu.Person = new(Person)
	stu.Name = "张三"
	fmt.Println(*stu.Person)

	tea := new(Teacher)
	tea.Name = "李四"
	fmt.Println((*tea).Person)

}

