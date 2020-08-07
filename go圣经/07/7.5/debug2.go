package main

import (
	"bytes"
	"io"
)

const debug2 = false

func main() {
	var buf io.Writer
	if debug2 {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f2(buf) // OK

}

// If out is non-nil, output will be written to it.
func f2(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}

}

