package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f,err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil { fmt.Println(err ) }

	counts := make(map[string]int )

	buf := bufio.NewScanner(f)
	buf.Split(bufio.ScanWords)
	for buf.Scan(){
		counts[buf.Text()] ++
	}

	fmt.Println(counts)
}
