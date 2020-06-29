package jsonreport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 只需要去除udid
type Report struct {
	Udid 		string 	`json:"udid"`
}


func ReportCount( countsUdid *map[string]int ) {
	var reporter Report
	var filename string
	lenargs := len(os.Args)
	if lenargs <= 1 {
		filename = "/tmp/jblog.log-20200618"
	}else {
		filename = os.Args[1]
		//filename = "/tmp/jblog.log-20200618"
	}
	//读文件
	f , err := os.Open(filename)
	if err != nil {
		fmt.Print(err )
	}
	defer f.Close()
	buf := bufio.NewScanner(f)

	for buf.Scan() {
		s := buf.Text()
		index := strings.Index(s,"readTree is ")
		if index != -1 {
			s = s[index+13:len(s)-1]
			//fmt.Println(s)
			err := json.NewDecoder(strings.NewReader(s)).Decode(&reporter)
			if err != nil {
				fmt.Println(err)
			}
			(*countsUdid)[reporter.Udid]++
		}
	}

}