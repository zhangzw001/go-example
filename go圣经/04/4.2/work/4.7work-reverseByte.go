package main

import "fmt"

//练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？


func reverseByte(s []byte) []byte {
	n := len(s)
	//for i,j := 0, n-1; i < j ; i,j = i +1, j-1 {
	for i,j := 0, n-1; i < n/2 ; i,j = i +1, j-1 {
		s[i] , s[j] = s[j] , s[i]
	}
	fmt.Printf("%p\n",&s)
	return s

}

func reverseBytePoint(s *[]byte) {
	n := len(*s)
	//for i,j := 0, n-1; i < j ; i,j = i +1, j-1 {
	for i,j := 0, n-1; i < n/2 ; i,j = i +1, j-1 {
		(*s)[i] , (*s)[j] = (*s)[j] , (*s)[i]
	}
	fmt.Printf("%p\n",s)
}
func main() {
	s := []byte{'1','2','3','4'}
	fmt.Printf("%p\n",&s)
	fmt.Printf("%c\n",reverseByte(s))
	fmt.Printf("%c\n",s)


	s2 := []byte{'1','2','3','4'}
	fmt.Printf("%p\n",&s2)
	reverseBytePoint(&s2)
	fmt.Printf("%c\n",s2)
}