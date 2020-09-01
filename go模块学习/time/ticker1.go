package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	for {
		<- ticker.C
		fmt.Println("...")
	}
}
