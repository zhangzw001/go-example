package main

import (
	"fmt"
)
type ByteCounter struct {

}
func main() {
	//var w io.Writer
	//w = os.Stdout

	//f := w.(*os.File)
	//c := w.(*byte.Buffer)

	//rw := w.(io.ReadWriter)
	//w = new(ByteCounter)


	s := "aa"
	//if v, ok := s.(string); ok {	//non-interface type string on left
	//	fmt.Println(v)
	//}
	if v, ok := interface{}(s).(string) ; ok {
		fmt.Println(v)
	}
}