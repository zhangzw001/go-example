package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)


func Charcount(){
	counts := make(map[rune]int)	//// counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	args := os.Args
	f , err := os.Open(args[1])
	defer f.Close()
	
	if err != nil {fmt.Println(err)}
	in := bufio.NewReader(f)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}


func Charcount2(f io.Reader) {
	counts := make(map[string]map[rune]int )
	//f,err := os.Open(os.Args[1])
	//defer f.Close()
	//if err != nil { fmt.Println(err )}
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