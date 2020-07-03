package fib

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 30 ; i++ {
			Fib(i)
		}
	}
}

func BenchmarkFib2(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 1000 ; i++ {
			Fib2(i)
		}
	}
}

func BenchmarkFib3(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 1000 ; i++ {
			Fib3()
		}
	}
}
