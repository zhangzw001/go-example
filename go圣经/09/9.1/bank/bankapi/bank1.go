package bankapi

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit1(amount int) { deposits <- amount }
func Balance1() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}