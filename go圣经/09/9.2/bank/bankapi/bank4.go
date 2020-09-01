package bankapi

import (
	"sync"
)

var (
	mu4 		sync.Mutex
	balance4 int
)

func Deposit4(amount int ) {
	mu4.Lock()
	defer mu4.Unlock()
	deposite4(amount)
}

func deposite4(amount int ) {
	balance4 += amount
}

func Balance4() int {
	mu4.Lock()
	defer mu4.Unlock()
	return balance4
}

func Withdraw4(amount int) bool  {
	mu4.Lock()
	defer mu4.Unlock()

	deposite4(-amount)

	if balance4 < 0 {
		//如果小于0,说明余额不足,还原
		deposite4(amount)
		return false
	}

	return true
}