package main

import (
	basename2 "github.com/zhangzw001/go-example/go圣经/03/3.5/3.5basename"
	"testing"
)


func BenchmarkBasename1(b *testing.B) {
	for i :=0 ; i < b.N ; i++ {
		s := "/var/log/cron.log"
		basename2.Basename1(s)
	}
}

func BenchmarkBasename2(b *testing.B) {
	for i :=0 ; i < b.N ; i++ {
		s := "/var/log/cron.log"
		basename2.Basename2(s)
	}
}


