package main

import (
	"encoding/json"
	"fmt"
	jsonreport "github.com/zhangzw001/go-example/go圣经/04/4.5/json-report"
	"strings"
)

func main() {
	s := `{
	"address":"上场",
	"udid":"c8da0f258b8e635714aad5adfsadf",
	"idfa":"00000000-0000-0000-0000-000000000000"
	}`
	var reporter jsonreport.Report

	json.NewDecoder(strings.NewReader(s)).Decode(&reporter)
	fmt.Println(reporter.Udid)

	//counts := make(map[string]int)
	//jsonreport.ReportCount(&counts)
	//fmt.Println(len(counts))
}
