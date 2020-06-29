package main
//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
import (
	"fmt"
	"unicode"
)

func delEmptyByte(s []byte ) []byte {
	len1 := len(s)
	for i:=0 ; i < len1 ; i++ {
		index := i+1
		if index >= len(s) {
			break
		}
		if unicode.IsSpace(rune(s[index])) && unicode.IsSpace(rune(s[i])) {
			//fmt.Println(i,index)
			copy(s[i:],s[index:])
			s=s[:len(s)-1]
			i--
		}
		//if s[index] != s[i] {
		//	s[n] = s[i]
		//	n++
		//}
	}
	return s[:]
}



func main() {
	s := []byte{' ','a',' ',' ',' ',' ','b'}
	s = delEmptyByte(s)
	fmt.Printf("%c,%d",s,len(s))

}
