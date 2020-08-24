package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	_ = 1 << (10 * iota)
	K
	M
	G
	T
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time
	n := make(chan struct{})
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				go func() {
					n <- struct{}{}
				}()
				break
				//break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case <-n:
			printDiskUsage(nfiles, nbytes)
			break loop
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	switch {
	case nbytes > G:
		fmt.Printf("%d files %.4fGB\n", nfiles, float64(nbytes)/float64(G))
	case nbytes > M:
		fmt.Printf("%d files %.4fMB\n", nfiles, float64(nbytes)/float64(M))
	case nbytes > K:
		fmt.Printf("%d files %.4fMB\n", nfiles, float64(nbytes/K))
	default:
		fmt.Printf("%d files %.4fB\n", nfiles, float64(nbytes))

	}
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			//fmt.Println(subdir, dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
