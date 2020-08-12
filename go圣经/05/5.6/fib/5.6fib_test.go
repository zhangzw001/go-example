package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	for i:= 0 ; i < 40 ; i++ {
		Fib(i)
	}
}

func TestFib2(t *testing.T) {
	for i:= 0 ; i < 90 ; i++ {
		Fib2(i)
	}
}

func TestFib3(t *testing.T) {
	f3 := Fib3()
	for i:= 0 ; i < 90 ; i++ {
		f3()
}
}

func BenchmarkFib(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 30 ; i++ {
			Fib(i)
		}
	}
}

func BenchmarkFib2(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 90 ; i++ {
			Fib2(i)
		}
	}
}

func BenchmarkFib3(b *testing.B) {
	f3 := Fib3()
	for i:=0 ; i < b.N ; i++ {
		for i:= 0 ; i < 90 ; i++ {
			f3()
		}
	}
}
