package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println(a)

	a2 := [...]int{1,3}
	fmt.Println(a2)

	a3 := [32]byte{9: '1'}
	zero(&a3)
	fmt.Println(a3)


}

func zero(p *[32]byte) {
	*p = [32]byte{}
	//for i := range p {
	//	p[i] = 0
	//}
}

