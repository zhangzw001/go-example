package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error ) {
	//(在这个+=赋值语句中，让len(p)的类型和*c的类型匹配的转换是必须的。)
	*c += ByteCounter(len(p))
	return len(p), nil

}

type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int , error) {
	var sum int
	buf := bufio.NewScanner(bytes.NewReader(p))
	buf.Split(bufio.ScanWords)
	for buf.Scan() {
		sum++
	}
	*w += WordsCounter(sum)
	return sum,nil
}

type LinesCounter int
func (l *LinesCounter) Write(p []byte) (int , error ) {
	var sum int
	buf := bufio.NewScanner(bytes.NewReader(p))
	buf.Split(bufio.ScanLines)
	for buf.Scan() {
		sum ++
	}
	*l += LinesCounter(sum )
	return sum,nil
}


//练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
//返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针
type CountWriter struct {
	Writer io.Writer
	Count int64
}

func (cw *CountWriter) Write(p []byte) (int , error ) {
	cw.Count += int
	return len(p), nil
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{
		Writer: w,
	}
	return &cw, &(cw.Count)
}

func main() {
	//7.1.1
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	str := "this is strings a a a "
	var w WordsCounter
	w.Write([]byte(str))
	fmt.Println(w)

	str2 := "a\nb\nc\nd\n"
	var l LinesCounter
	l.Write([]byte(str2))
	fmt.Println(l)
}