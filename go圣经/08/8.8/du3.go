package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
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

var sema = make(chan struct{}, 10)
func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()


	var tick <-chan time.Time
	n := make(chan struct{})
	if *verbose {
		tick = time.Tick(100 * time.Millisecond)
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

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			//fmt.Println(subdir, dir, entry.Name())
			go walkDir(subdir,wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
