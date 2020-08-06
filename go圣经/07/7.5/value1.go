package main

import (
	"fmt"
)

type A struct {
	string
}
func (a A) String() string {
	return fmt.Sprintf("this is String() : %v\n",a.string)
}
func main() {
	a := A{"test"}
	fmt.Println(a)
	//var w io.Writer
	//fmt.Println(w)
	//w = os.Stdout
	//w.Write([]byte("hello"))

}