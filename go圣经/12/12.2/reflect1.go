package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)


type TestRead struct {
	a int
}

func (t *TestRead) Read(b []byte) (n int, err error) {
	return t.a,nil
}

func (t *TestRead) Write(b []byte) (n int, err error) {
	return t.a,nil
}

func main() {
	//
	var a int = 3
	fmt.Println(reflect.TypeOf(a))

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	var r1 io.Reader = &TestRead{}
	fmt.Println(reflect.TypeOf(r1))

	var w1 io.Writer = &TestRead{}
	fmt.Println(reflect.TypeOf(w1))
	//Cannot use '&TestRead{}' (type *TestRead) as type io.Writer in assignment Type does not implement 'io.Writer'
	//as some methods are missing: Write(p []byte) (n int, err error)

}
