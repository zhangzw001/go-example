package main

import (
	"fmt"
	"math/rand"
)

var heads,tails int

func main() {

	heads,tails = Coinflip(1000)
	fmt.Println(heads,tails)

}

func Coinflip(n int ) (int, int ){
	//var heads,tails int
	for i:=0;i<n;i++ {
		switch coinflip() {
		case 0:
			heads++
		case 1:
			tails++
		default:
			fmt.Println("landed on edge!")
		}
	}
	return heads,tails
}

func coinflip() int {
	return rand.Intn(2)
}
