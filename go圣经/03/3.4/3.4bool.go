package main

import "fmt"

func BToI( b bool ) int {
	if b {
		return 1
	}
	return 0
}
// 1 => true , 0 => false
func IToB( i int ) bool {
	return i != 0
}

func main () {
	b1 := true
	fmt.Println(BToI(b1))

	for i := 0 ; i < 3 ; i ++ {
		fmt.Println(IToB(i))
	}
}

