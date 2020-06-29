package main

import (
	"fmt"
	tempconv0 "github.com/zhangzw001/go-example/go圣经/02/2.5"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		//fmt.Printf("%p\n",arg)
		t, err := strconv.ParseFloat(arg, 64)
		//fmt.Printf("%p\n",t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv0.Fahrenheit(t)
		c := tempconv0.Celsius(t)
		//fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		//	f, tempconv0.FToC(f), c, tempconv0.CToF(c), c, tempconv0.CToK(c))
		fmt.Printf("华氏:%f = 摄氏:%f\n摄氏:%f = 华氏:%f\n摄氏:%f = 开氏:%f\n",
			f, tempconv0.FToC(f), c, tempconv0.CToF(c), c, tempconv0.CToK(c))
	}
}