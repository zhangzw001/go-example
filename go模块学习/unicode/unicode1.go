package main

import (
	"fmt"
	"unicode"
)

const (
	pC     = 1 << iota // a control character.
	pP                 // a punctuation character.
	pN                 // a numeral.
	pS                 // a symbolic character.
	pZ                 // a spacing character.
	pLu                // an upper-case letter.
	pLl                // a lower-case letter.
	pp                 // a printable character according to Go's definition.
	pg     = pp | pZ   // a graphical character according to the Unicode definition.
	pLo    = pLl | pLu // a letter that is neither upper nor lower case.
	pLmask = pLo
)

func main() {
	fmt.Println(pZ,pp,pg)
	fmt.Printf("%b,%b",128,144)
	s := "天气,下雨 zzz ZZZ"
	for _, r := range s {
		if unicode.IsLetter(r) {
			fmt.Println(string(unicode.ToLower(r)))
		}else {
			fmt.Println(string(r))
		}
	}
}
