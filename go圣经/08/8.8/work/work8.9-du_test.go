package dudu

import (
	"testing"
)

func BenchmarkDu(t *testing.B) {
	for i :=0 ; i < t.N ; i ++ {
		Du([]string{"/Users/zhangzw/work"})
	}
}

//func BenchmarkFileSize(t *testing.B) {
//	for i :=0 ; i < t.N ; i ++ {
//		tmp.sumFileSize([]string{"/Users/zhangzw/work"})
//	}
//}
