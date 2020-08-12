package main

import (
	"errors"
	"fmt"
	"reflect"
)
type p struct {
	s string
}
func main() {
	//ok := errors.New("EOF")
	//fmt.Println(ok,reflect.TypeOf(ok),ok.Error(),reflect.TypeOf(ok.Error()))
	//fmt.Println(errors.New("EOF").Error() == errors.New("EOF").Error())

	//fmt.Printf("%p,%p\n",errors.New("EOF"),errors.New("EOF"))
	//fmt.Println(reflect.TypeOf(errors.New("EOF")),reflect.ValueOf(errors.New("EOF")))
	//fmt.Println(reflect.TypeOf(errors.New("EOF")),reflect.TypeOf(errors.New("EOF")))
	fmt.Println(errors.New("EOF") == errors.New("EOF")) //false
	p1 := &p{"EOF"}
	p2 := &p{"EOF"}
	fmt.Println(reflect.TypeOf(p1),reflect.TypeOf(p2))
	fmt.Printf("%v,%[2]v,%[1]p,%[2]p\n",p1,&p1)
	fmt.Printf("%v,%[2]v,%[1]p,%[2]p\n",p2,&p2)
	fmt.Println(p1 == p2 )
	p3 := &p{}
	p4 := &p{}
	fmt.Println(p3 == p4 )

	//var p *int
	//a := 1
	//p = &a


}
