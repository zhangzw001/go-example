package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(func(r rune) rune{ return r + 1 }, "HAL-9000"))
}
