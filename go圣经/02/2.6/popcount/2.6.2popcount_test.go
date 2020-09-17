package popcount

import (
	"reflect"
	"testing"
)

func assert(t *testing.T,expected,actual interface{}){
	if !reflect.DeepEqual(expected,actual){
		t.Errorf("(expected,actual) = (%v,%v)\n",expected,actual)
	}
}

func ExamplePopCount() {
	PopCount(0x1234567890ABCDEF)
}

func TestPopCount(t *testing.T) {
	assert(t,32, PopCount(0x1234567890ABCDEF))
}


func TestPopCountByLoop(t *testing.T) {
	assert(t, 32, PopCountByLoop(0x1234567890ABCDEF))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}
func BenchmarkPopCountByLoopShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoopShift(0x1234567890ABCDEF)
	}
}
func BenchmarkPopCountByAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByAnd(0x1234567890ABCDEF)
	}
}

