package main

import (
	"fmt"
	"math"
)

func main() {

	const Avogadro = 6.02214129e23  // 阿伏伽德罗常数
	const Planck   = 6.62606957e-34 // 普朗克常数
	fmt.Printf("%f %[1]g %[1]e \n%[2]f %[2]g %[2]e \n",Avogadro,Planck)

	for x := 0; x < 16; x++ {
		// 宽度12,精度3位小数
		fmt.Printf("x = %2d e^x = %12.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}