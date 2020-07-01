package tmp

import "testing"

func BenchmarkCountWordsAndImages(b *testing.B) {
	for i:=0 ; i < b.N ;i ++ {
		cWaI("https://www.baidu.com")
	}
}
