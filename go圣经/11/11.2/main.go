package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	fmt.Printf("Random seed: %d",seed)
	r := rand.New(rand.NewSource(seed))

}
