package main

import (
	"log"
	"os"
)

func main() {

	// 下面这种会导致 无法关闭, 在for循环结束之后, 可能会导致文件描述符耗尽
	filenames := os.Args[1:]
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close() // NOTE: risky; could run out of file
		//descriptors
		// ...process f…
	}

	// 一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			log.Fatal(err)
		}
	}
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f…
	return err
}
