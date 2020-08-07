package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

const debug2 = false

func main() {
	var buf io.Writer
	fmt.Printf("%T\n",buf)
	if debug2 {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f2(buf) // OK

}

// If out is non-nil, output will be written to it.
func f2(out io.Writer) {
	// ...do something...
	if out != nil {
		fmt.Println(reflect.TypeOf(out),reflect.ValueOf(out))
		out.Write([]byte("done!\n"))
	}

}

