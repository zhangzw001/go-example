package bankapi



var balance int
func Deposit(amount int) { balance = balance + amount }
func Balance() int { return balance }