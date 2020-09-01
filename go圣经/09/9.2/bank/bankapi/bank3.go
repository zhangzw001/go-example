package bankapi

import (
	"sync"
)

var (
	mu 		sync.Mutex
	balance3 int
)

func Deposit3(amount int ) {
	mu.Lock()
	balance3 = balance3 + amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	return balance3
}

func Withdraw3(amount int) error {
	if Balance3() >= amount {
		mu.Lock()
		balance3 = balance3 - amount
		mu.Unlock()
	}else {
		//return fmt.Errorf("余额不足")
	}
	return nil
}