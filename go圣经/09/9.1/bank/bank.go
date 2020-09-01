package main

import (
	"bank/bankapi"
	"fmt"
	"time"
)

func main() {
	go func() {
		//bankapi.Deposit(200)
		//fmt.Println("=",bankapi.Balance())
		bankapi.Deposit1(200)
		fmt.Println("=",bankapi.Balance1())
	}()

	go func() {
		//bankapi.Deposit(100)
		//fmt.Println("=",bankapi.Balance())
		bankapi.Deposit1(100)
		fmt.Println("=",bankapi.Balance1())
	}()
	time.Sleep(1 * time.Second)
}
