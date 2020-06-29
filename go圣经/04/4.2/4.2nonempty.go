package main

import "fmt"

func nonempty(s []string) []string {
	i:=0
	for _,j := range s {
		if j != "" {
			s[i] = j
			i++
		}
	}
	return s[:i]
}

func nonemptyAppend(s []string) []string {
	rst := s[:0]
	for _, j := range s {
		if j != "" {
			rst = append(rst, j)
		}
	}
	return rst
}

func main() {
	// nonempty 测试
	s1 := []string{"1","","3","","5"}
	// 以下由于 slice共享同一个底层数组 , 所以s1也被改变
	fmt.Println(nonempty(s1))		//[1 3 5]
	fmt.Println(s1)					//[1 3 5  5]
	// 因此我们一边按照以下方式调用
	s2 := []string{"1","","3","","5"}
	s2 = nonempty(s2)
	fmt.Println(s2)

	// nonemptyAppend 测试
	s3 :=[]string{"1","","3","","5"}
	s3 = nonemptyAppend(s3)
	fmt.Println(s3)
}