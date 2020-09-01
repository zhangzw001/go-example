package main

import (
	"bank/bankapi"
	"fmt"
	"time"
)

func main() {
	for i:=0 ; i < 10 ; i++ {
		go func() {
			bankapi.Deposit3(200)
			fmt.Println("存入200=", bankapi.Balance3())
		}()

		go func() {
			bankapi.Deposit3(100)
			fmt.Println("存入100=", bankapi.Balance3())
		}()

		go func() {

			err := bankapi.Withdraw3(300)
			if err != nil {
				fmt.Println(err)
			}else {
				fmt.Println("成功取款199后=", bankapi.Balance3())
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("最后余额=", bankapi.Balance3())

}
