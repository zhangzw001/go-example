package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	p1 := new(int)
	fmt.Println(p,p1,p==p1)

	fmt.Println(delta(100 , 10))


}


func delta(old, new int ) int {return new - old }