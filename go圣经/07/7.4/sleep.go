package main

import (
	"flag"
	"time"
)

var period = flag.Duration("p",1*time.Second,"sleep period")

func main() {
	flag.Parse()
	
}