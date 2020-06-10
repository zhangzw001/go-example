package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()
	counts := make(map[string]int)

	//files := os.Args[1:]

	files := []string{"./go圣经/01/1.3/1.3bufio-learn1.txt"}

	if len(files) == 0 {
		countLines(os.Stdin,counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}

	for line, n := range counts {
			fmt.Printf("%d\t\t%s\n",n, line)
	}

	fmt.Printf("%.6f",time.Since(start).Seconds())
}


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}