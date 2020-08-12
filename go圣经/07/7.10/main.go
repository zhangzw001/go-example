package main

import (
	"io"
	"os"
)
type ByteCounter struct {

}
func main() {
	var w io.Writer
	w = os.Stdout

	//f := w.(*os.File)
	//c := w.(*byte.Buffer)

	rw := w.(io.ReadWriter)
	w = new(ByteCounter)
}