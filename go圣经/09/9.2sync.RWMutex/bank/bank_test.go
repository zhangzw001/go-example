package main

import (
	"bank/bankapi"
	"testing"
)
func BenchmarkBank(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit2(200)
			bankapi.Balance2()
		}()
		go func() {
			bankapi.Deposit2(100)
			bankapi.Balance2()
		}()
		go func() {
			bankapi.Withdraw2(300)
			bankapi.Balance2()
		}()
	}
}


func BenchmarkBank3(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit3(200)
			bankapi.Balance3()
		}()
		go func() {
			bankapi.Deposit3(100)
			bankapi.Balance3()
		}()
		go func() {
			bankapi.Withdraw3(300)
			bankapi.Balance3()
		}()
	}
}

func BenchmarkBank4(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit4(200)
			bankapi.Balance4()
		}()
		go func() {
			//bankapi.Deposit4(100)
			bankapi.Balance4()
		}()
		go func() {
			//bankapi.Withdraw4(300)
			bankapi.Balance4()
			bankapi.Balance4()
			bankapi.Balance4()
			bankapi.Balance4()
			bankapi.Balance4()
		}()
	}
}

func BenchmarkBank5(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit5(200)
			bankapi.Balance5()
		}()
		go func() {
			//bankapi.Deposit5(100)
			bankapi.Balance5()
		}()
		go func() {
			//bankapi.Withdraw5(300)
			bankapi.Balance5()
			bankapi.Balance5()
			bankapi.Balance5()
			bankapi.Balance5()
			bankapi.Balance5()
		}()
	}
}