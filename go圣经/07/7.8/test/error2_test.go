package test

import "testing"

func BenchmarkError2(b *testing.B) {
	for i := 0 ; i < b.N ; i ++ {
		error2()

	}
}
