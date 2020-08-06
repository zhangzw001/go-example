package main

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	AbsoluteZeroC 	Celsius = -273.15	//绝对零度
	FreezingC 		Celsius = 0			//冰点
	BoilingC      	Celsius = 100     	// 沸水温度

)
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64		// 开氏温度

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kelvin) Celsius { return Celsius(k)-AbsoluteZeroC}
type celsiusFlag struct {
	Celsius
}
func (f *celsiusFlag) String() (string) {
	return strconv.Itoa(int(f.Celsius))
}

func (f *celsiusFlag) Set(s string ) error {
	var unit string
	var value float64
	//调用fmt.Sscanf函数从输入s中解析一个浮点数（value）和一个字符串（unit）。虽然通常必须检查Sscanf的错误返回，但是在这个例子中我们不需要因为如果有错误发生，就没有switch case会匹配到。
	fmt.Sscanf(s,"%f%s",&value,&unit)
	//fmt.Println(len("%f%s"))
	//fmt.Println(value,unit)
	switch unit {
	case "C","°C":
		f.Celsius = Celsius(value)
		return nil
	case "F","°F":
		f.Celsius = FToC(Fahrenheit(value))
		fmt.Println(f.Celsius)
		return nil
	case "K","°K":
		f.Celsius = KToC(Kelvin(value))
	}
	return fmt.Errorf("invalid temperature %q",s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
//func CelsiusFlag(name string, value Celsius, usage string ) *Celsius {
//	f := celsiusFlag{value}
//	flag.CommandLine.Var(&f , name , usage)
//	return &f.Celsius
//}
func main() {
	//var c celsiusFlag
	//c.Set("100F")
	//fmt.Println(c)
	//c.Set("100K")
	//fmt.Println(c)
	var temp = CelsiusFlag("temp",20.0,"the temperature")
	flag.Parse()
	fmt.Println(*temp)
}