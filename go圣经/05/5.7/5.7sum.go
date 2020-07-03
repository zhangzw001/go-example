package main

import (
	"fmt"
	"os"
)



func main() {
	fmt.Println(sum(1,2,3,4))

	vars := []int{1,2,3,4}
	fmt.Println(sum(vars...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"

	fmt.Println(sum())
}

func sum(vals ...int) int {
	total := 0
	if len(vals) == 0 {
		errorf(24,"vals is nil : %v",vals)
		return total

	}
	for _, val := range vals {
		total += val
	}
	return total
}


func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

