package main

import (
	"github.com/zhangzw001/go-example/go圣经/02/2.6/popcount"
	"reflect"
	"testing"
)

func assert(t *testing.T,expected,actual interface{}){
	if !reflect.DeepEqual(expected,actual){
		t.Errorf("(expected,actual) = (%v,%v)\n",expected,actual)
	}
}

func TestPopCount(t *testing.T) {
	assert(t,32,popcount.PopCount(0x1234567890ABCDEF))
}


func TestPopCountByLoop(t *testing.T) {
	assert(t, 32, popcount.PopCountByLoop(0x1234567890ABCDEF))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(0x1234567890ABCDEF)
	}
}
func BenchmarkPopCountByLoopShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoopShift(0x1234567890ABCDEF)
	}
}
func BenchmarkPopCountByAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByAnd(0x1234567890ABCDEF)
	}
}

