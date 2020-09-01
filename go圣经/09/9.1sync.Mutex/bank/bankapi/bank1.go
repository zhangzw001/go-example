package bankapi

import (
	"fmt"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan int)

func Deposit1(amount int) { deposits <- amount }
func Balance1() int       { return <-balances }

//练习 9.1sync.Mutex： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
//其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。这条消息会被发送给monitor的goroutine，
//且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
func Withdraw(amount int) error {
	m := <- balances
	if m >= amount {
		withdraw <- amount
	}else {
		return fmt.Errorf("余额不足")
	}
	return nil
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraw:
			balance -= amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
