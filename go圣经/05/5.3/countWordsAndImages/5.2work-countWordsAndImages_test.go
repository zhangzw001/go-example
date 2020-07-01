package countWordsAndImages

import (
	"log"
	"testing"
)

func BenchmarkCountWordsAndImages(b *testing.B) {
	for i:=0 ; i < b.N ;i ++ {
		_,_,err := CountWordsAndImages("https://www.baidu.com")
		if err != nil { log.Fatal(err)}
	}
}
