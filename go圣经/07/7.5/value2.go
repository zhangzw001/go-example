package main

import (
	"fmt"
)

type ss interface {
	speak(string) string
}

type stu struct{}
func (s stu) speak(str string) string{
	return str
}
func live() ss{
	var stu stu
	return stu
}

func TestLive(){
	if live() == nil {            //not nil
	//if reflect.ValueOf(live()).IsNil(){           //panic
		fmt.Println("AAAAAAAA")
	}else {
		fmt.Println("BBBBBBBB")
	}
}


//2

func main() {
	//TestLive()

}

