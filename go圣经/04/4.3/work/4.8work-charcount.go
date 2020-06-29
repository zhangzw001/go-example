package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	charcount2()
}


func charcount2() {
	counts := make(map[string]map[rune]int )
	f,err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil { fmt.Println(err )}

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		buf.Bytes()
		for _,v := range []rune(buf.Text()) {
			if counts["letter"] == nil {
				counts["letter"] = make(map[rune]int)
			}
			if counts["number"] == nil {
				counts["number"] = make(map[rune]int)
			}
			if unicode.IsLetter(v) {
				counts["letter"][v]++
			}else if unicode.IsNumber(v) {
				counts["number"][v]++
			}
		}

	}

	for k,v := range counts{
		for i,j := range v {
			if k == "letter" {
				fmt.Printf("字符: %s, %d\n",string(i), j)
			}
			if k == "number" {
				fmt.Printf("数字: %d, %d\n",int(i),j)
			}
		}
	}

}