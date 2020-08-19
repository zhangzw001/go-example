package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CopyF struct {
	srcFileName string
	dstFileName string
}



func (c CopyF) copyFile() {
	fsrc , err := os.Open(c.srcFileName)
	defer fsrc.Close()
	if err != nil { log.Fatalf("源文件打开失败: %v\n",err )}
	fdst , err := os.OpenFile(c.dstFileName,os.O_CREATE|os.O_WRONLY,0644)
	defer fdst.Close()
	if err != nil { log.Fatalf("目标文件打开失败: %v\n",err )}

	buf := bufio.NewScanner(fsrc)
	for buf.Scan() {
		str := fmt.Sprintf("%v\n",buf.Text())
		str = strings.Trim(str," ")
		fdst.WriteString(str)
	}
}

func main() {
	copyfile := &CopyF{"file1.txt","file1-bak.txt"}
	copyfile.copyFile()
}
