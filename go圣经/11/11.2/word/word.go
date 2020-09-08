package word

import (
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
	return true
}
