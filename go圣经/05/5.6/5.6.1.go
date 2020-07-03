package main

import (
	"fmt"
	"os"
)

func main() {
	var rmdirs []func()
	for _, d := range os.TempDir() {
		dir := fmt.Sprintf("./test/%s",string(d))
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})

	}
	fmt.Println(rmdirs)
	//for _, rmdir := range rmdirs {
	//	rmdir()
	//}
}