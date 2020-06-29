package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Meter float64
type Feet float64
type Pound float64
type Kilogram float64


//许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法
func (M Meter) String() string {return fmt.Sprintf("%gm",M)}
func (F Feet) String() string { return fmt.Sprintf("%gft",F)}
func (P Pound) String() string {return fmt.Sprintf("%glb",P)}
func (K Kilogram) String() string {return fmt.Sprintf("%gkg",K)}

func MToF(M Meter) Feet {return Feet(M * 1200 / 3937)}
func FToM(F Feet) Meter {return Meter(F * 3937 / 1200)}
func PToK(P Pound) Kilogram {return Kilogram( P * 0.45359237)}
func KToP(K Kilogram) Pound { return Pound(K / 0.45359237)}

var flagMeter = flag.Float64("m",0,"Meter  float64")
var flagFeet = flag.Float64("f",0,"Feet  float64")
var flagPound = flag.Float64("p",0,"Pound  float64")
var flagKilogram = flag.Float64("k",0,"Kilogram  float64")


func main() {
	//flag.Parse()
	//fmt.Printf("%p\n",Meter(*flagMeter))
	//m1 :=Meter(*flagMeter)
	//fmt.Println(MToF(m1))
	//var args []float64{}
	if len(os.Args) > 1 {
		flag.Parse()
		m1 := Meter(*flagMeter)
		f1 := Feet(*flagFeet)
		p1 := Pound(*flagPound)
		k1 := Kilogram(*flagKilogram)
		//fmt.Printf("%f,%f,%f,%f\n",m1,f1,p1,k1)
		if m1 != 0 {fmt.Println(MToF(m1))}
		if f1 != 0 {fmt.Println(FToM(f1))}
		if p1 != 0 {fmt.Println(PToK(p1))}
		if k1 != 0 {fmt.Println(KToP(k1))}
	}else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(">>> 请输入转换的类型 <m|f|p|k> : ")
		fmt.Printf(">>> ")
		for scanner.Scan(){
			args := scanner.Text()
			switch args {
			case "m":
					fmt.Println(">>> 转换的是:Meter")
					fmt.Printf(">>> 请输入一个Meter值: ")
					scanner.Scan()
					tmp,_ := strconv.ParseFloat(scanner.Text(),64)
					m1 := Meter(tmp)
					fmt.Printf("%s = %s\n",m1,MToF(m1))
					fmt.Printf(">>> ")
			case "f":
					fmt.Println(">>> 转换的是:Feet")
					fmt.Printf(">>> 请输入一个Feet值: ")
					scanner.Scan()
					tmp,_ := strconv.ParseFloat(scanner.Text(),64)
					f1 := Feet(tmp)
					fmt.Printf("%s = %s\n",f1,FToM(f1))
					fmt.Printf(">>> ")
			case "p":
					fmt.Println(">>> 转换的是:Pound")
					fmt.Printf(">>> 请输入一个Pound值: ")
					scanner.Scan()
					tmp,_ := strconv.ParseFloat(scanner.Text(),64)
					p1 := Pound(tmp)
					fmt.Printf("%s = %s\n",p1,PToK(p1))
					fmt.Printf(">>> ")

			case "k":
					fmt.Println(">>> 转换的是:Kilogram")
					fmt.Printf(">>> 请输入一个Kilogram值: ")
					scanner.Scan()
					tmp,_ := strconv.ParseFloat(scanner.Text(),64)
					k1 := Kilogram(tmp)
					fmt.Printf("%s = %s\n",k1,KToP(k1))
					fmt.Printf(">>> ")

			case "help": fmt.Println(">>> 请输入转换的类型 <m|f|p|k> : ")
			default : fmt.Printf(">>> ")
			}

		}

	}
}

