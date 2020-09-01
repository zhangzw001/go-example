package bankapi

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance2 int
)

func Deposit2(amount int) {
	sema <- struct{}{} // acquire token
	balance2 = balance2 + amount
	<-sema // release token
}

func Balance2() int {
	sema <- struct{}{} // acquire token
	b := balance2
	<-sema // release token
	return b
}

func Withdraw2(amount int) error {
	if Balance2() >= amount {
		sema <- struct{}{}
		balance2 = balance2 - amount
		<- sema
	}else {
		//return fmt.Errorf("余额不足")
	}
	return nil
}