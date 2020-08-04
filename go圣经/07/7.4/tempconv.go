package main

import (
	"fmt"
	tempconv0 "github.com/zhangzw001/go-example/go圣经/02/2.5"
)


type celsiusFlag struct {
	tempconv0.Celsius
}

func (f *celsiusFlag) Set(s string ) error {
	var unit string
	var value float64
	fmt.Sscanf(s,"%f%s",&value,&unit)
	fmt.Println(value,unit)
	switch unit {
	case "C","°C":
		f.Celsius = tempconv0.Celsius(value)
		return nil
	case "F","°F":
		f.Celsius = tempconv0.FToC(tempconv0.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q",s)
}

func main() {
	var c celsiusFlag
	c.Set("100F")
	fmt.Println(c)
}