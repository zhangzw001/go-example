package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		
	}
}

func walkDir(dir string, fileSizes chan<-int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir,fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}


func dirents(dir string) []os.FileInfo {
	entries,err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr,"du1: ^v\n",err)
		return nil
	}
	return entries
}