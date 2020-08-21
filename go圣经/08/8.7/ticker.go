package main

import "time"

func main() {
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C
	ticker.Stop()
}
