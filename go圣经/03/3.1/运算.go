package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 35
	var n2 int = 40
	var avg float64
	avg = float64( n1 + n2 ) / 2
	fmt.Println(avg)



	var mod float64
	mod = float64( n1 / n2 )
	fmt.Println(mod)

	var n3 int = 2
	avg = float64(n1+n2) / float64(n3)
	fmt.Println(avg)


	var b float64
	var c int64
	a := "3.14"
	a1 := "3"
	//a = "111"
	//b = int(a)
	//b, _ = strconv.Atoi(a)
	c, _ = strconv.ParseInt(a1,10,64)
	b , _ = strconv.ParseFloat(a,10)
	fmt.Println(a,b,c)

	// 1
	//buf := bufio.NewReader(os.Stdin)
	//r,_ := buf.ReadString('\n')
	//r = strings.ReplaceAll(r,"\n","")
	////fmt.Println(r)
	//year,_ := strconv.Atoi(r)
	////fmt.Println(year)
	// 2
	var year int
	fmt.Scanf("%d",&year)
	fmt.Println(year)
	if (year % 400 ) == 0 || ( (year % 4) == 0 && (year % 100) != 0 ) {
		fmt.Println("闰年")
	}


}
