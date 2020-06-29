package main

import (
	jsonreport "github.com/zhangzw001/go-example/go圣经/04/4.5/json-report"
	"testing"
)

func BenchmarkReportCount(b *testing.B) {
	counts := make(map[string]int)
	for i := 0 ; i < b.N ; i ++ {
		jsonreport.ReportCount(&counts)
	}
}
