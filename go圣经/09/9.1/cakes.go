package main

import (
	"fmt"
	"time"
)

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		fmt.Println("baker")
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked2 <-chan *Cake) {
	for cake := range cooked2 {
		cake.state = "iced"
		fmt.Println("icer")
		iced <- cake // icer never touches this cake again
	}
}
func main() {
	cake1 := make(chan *Cake)
	//timeout := make(chan time.Time)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	timeout <-
	//}()
		ticker := time.NewTicker(2 * time.Second)

		go baker(cake1)
		loop:
		for {
			select {
			case <-ticker.C:
				fmt.Println("退出...")
				ticker.Stop()
				break loop
			default:
				go baker(cake1)
				go icer(cake1,cake1)
				time.Sleep(1 * time.Second)
			}
		}
}
