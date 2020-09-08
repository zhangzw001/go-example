package _for

import "testing"

func BenchmarkTest1(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		test1()
	}
}


func BenchmarkTest2(t *testing.B) {
	for i := 0 ; i < t.N ; i ++ {
		test2()
	}
}