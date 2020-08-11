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
	fmt.Println(errors.New("EOF") == errors.New("EOF"))
	fmt.Printf("%p,%p\n",errors.New("EOF"),errors.New("EOF"))

	p1 := new(p)
	p2 := new(p)
	fmt.Printf("%p,%p\n",p1,p2)
}
