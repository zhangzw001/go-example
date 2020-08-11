package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(errors.New("EOF"))
	fmt.Println(errors.New("EOF") == errors.New("EOF"))
}
