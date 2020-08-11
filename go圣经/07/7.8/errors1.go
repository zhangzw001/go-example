package main

import (
	"errors"
	"fmt"
)
type p struct {
	//s string
}
func main() {
	//ok := errors.New("EOF")
	//fmt.Println(ok,reflect.TypeOf(ok),ok.Error(),reflect.TypeOf(ok.Error()))
	//fmt.Println(errors.New("EOF").Error() == errors.New("EOF").Error())
	fmt.Println(errors.New("EOF") == errors.New("EOF")) //false
	//
	//fmt.Printf("%p,%p\n",errors.New("EOF"),errors.New("EOF"))
	//
	//fmt.Println(reflect.TypeOf(errors.New("EOF")),reflect.ValueOf(errors.New("EOF")))
	//fmt.Println(reflect.TypeOf(errors.New("EOF")),reflect.TypeOf(errors.New("EOF")))
	p1 := &p{}
	p2 := &p{}
	//fmt.Println(reflect.TypeOf(p1),reflect.TypeOf(p2))
	fmt.Printf("%p,%v\n",p1,&(p{}))
	fmt.Println(p1,p2,&p1,&p2)
	fmt.Printf("%p,%v,%p\n",p1,&p1,&p1)
	if p1 == p2 {
		fmt.Println("相等")
	}else {
		fmt.Println("不等")
	}

	var p *int
	a := 1
	p = &a
	fmt.Printf("%p,%v,%p\n",p,&a,&a)


}
