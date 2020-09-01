package bankapi

import (
	"sync"
)

var (
	mu5 		sync.RWMutex
	balance5 int
)

func Deposit5(amount int ) {
	mu5.Lock()
	defer mu5.Unlock()
	deposite5(amount)
}

func deposite5(amount int ) {
	balance5 += amount
}

func Balance5() int {
	mu5.RLock()
	defer mu5.RUnlock()
	return balance5
}

func Withdraw5(amount int) bool  {
	mu5.Lock()
	defer mu5.Unlock()

	deposite5(-amount)

	if balance5 < 0 {
		// 如果小于0,说明余额不足,还原
		// 但如果某个时刻刚好程序在还原前挂了, 会导致数据异常吧?
		deposite5(amount)
		return false
	}

	return true
}