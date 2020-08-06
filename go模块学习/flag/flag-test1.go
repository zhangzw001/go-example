package main

import (
	"flag"
	"fmt"
)

func main() {
	//1
	//var flagtest1 = flag.Float64("f",100.0,"test float64")
	var flagtest1 = flag.CommandLine.Float64("f",100.0,"test float64")
	flag.Parse()
	switch {
	case *flagtest1 > 100 :
		fmt.Println(">")
	case *flagtest1 < 100 :
		fmt.Println("<")
	default:
		fmt.Println("=")
	}

}
