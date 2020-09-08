package word

import (
	"fmt"
	"math/rand"
	"unicode"
)

func IsPalindrome(s string ) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}


//我们现在的任务就是修复这些错误。简要分析后发现第一个BUG的原因是我们采用了 byte而不是rune序列，所以像“été”中的é等非ASCII字符不能正确处理。第二个BUG是因为没有忽略空格和字母的大小写导致的。
func IsPalindrome2(s string) bool {
	var runes []rune
	for _,i := range s {
		if unicode.IsLetter(i) {
			runes = append(runes, unicode.ToLower(i))
		}
	}
	for i := range runes {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	fmt.Println("true: ",string(runes))
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func RandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func RandomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(20)+5 // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		//r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		r := rune(rng.Intn(128))
		runes[i] = r
	}
	return string(runes)
}
