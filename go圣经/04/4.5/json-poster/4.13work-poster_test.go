package main

import (
	work_poster "github.com/zhangzw001/go-example/go圣经/04/4.5/json-poster/src"
	"testing"
)

func BenchmarkGetPoster(b *testing.B) {
	for i := 0 ;i < b.N ; i ++ {
		work_poster.GetPoster("Holmes")
	}
}

func BenchmarkGetPosterSpritf(b *testing.B) {
	for i := 0 ; i < b.N ;i ++ {
		work_poster.GetPosterSprintf("Holmes")
	}
}

func BenchmarkGetPosterAdd(b *testing.B) {
	for i := 0 ; i < b.N ;i ++ {
		work_poster.GetPosterAdd("Holmes")
	}
}
