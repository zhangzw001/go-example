package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	//动态分配机制依然决定(*bytes.Buffer).Write的方法会被调用，但是这次的接收者的值是nil。
	//对于一些如*os.File的类型，nil是一个有效的接收者(§6.2.1)，但是*bytes.Buffer类型不在这些类型中。这个方法会被调用，但是当它尝试去获取缓冲区时会发生panic。
	if out != nil {
		out.Write([]byte("done!\n"))
	}

}

