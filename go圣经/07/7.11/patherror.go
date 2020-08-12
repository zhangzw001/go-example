package main

import (
	"fmt"
	"os"
)

//// PathError records an error and the operation and file path that caused it.
//type PathError struct {
//	Op   string
//	Path string
//	Err  error
//}
//
//func (e *PathError) Error() string {
//	return e.Op + " " + e.Path + ": " + e.Err.Error()
//}

var ErrNotExist
func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n",err)
}
