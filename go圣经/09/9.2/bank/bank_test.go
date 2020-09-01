package main

import (
	"bank/bankapi"
	"testing"
)
func BenchmarkBank(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit2(200)
		}()
		go func() {
			bankapi.Deposit2(100)
		}()
		go func() {
			bankapi.Withdraw2(300)
		}()
	}
}


func BenchmarkBank3(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit3(200)
		}()
		go func() {
			bankapi.Deposit3(100)
		}()
		go func() {
			bankapi.Withdraw3(300)
		}()
	}
}

func BenchmarkBank4(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		go func() {
			bankapi.Deposit4(200)
		}()
		go func() {
			bankapi.Deposit4(100)
		}()
		go func() {
			bankapi.Withdraw4(300)
		}()
	}
}
