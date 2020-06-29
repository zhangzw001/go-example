package tempconv0

import "fmt"

const (
	AbsoluteZeroC 	Celsius = -273.15	//绝对零度
	FreezingC 		Celsius = 0			//冰点
	BoilingC      	Celsius = 100     	// 沸水温度

)

// 命名类型, 为的是显示操作, 例如 Celsius(10) 是将10转换为Celsius类型, 也就是float64类型
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64		// 开氏温度

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin {
	return Kelvin(c + AbsoluteZeroC)
}
func KToC(k  Kelvin) Celsius {
	return Celsius(k) - AbsoluteZeroC
}
func (c Celsius) String()  { fmt.Printf("%g°C", c) }

//func main() {
//	fmt.Println(CToF(100))
//	fmt.Println(FToC(100))
//
//	//
//	c := FToC(100)
//	fmt.Printf("%p\n",c)
//	f := CToF(100)
//	fmt.Printf("%p\n",f)
//	c.String()
//
//}

