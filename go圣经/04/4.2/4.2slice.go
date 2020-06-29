package main

import "fmt"

func reverse(s []string )  {
	n := len(s)
	//for i,j := 0, n-1; i < j ; i,j = i +1, j-1 {
	for i,j := 0, n-1; i < n/2 ; i,j = i +1, j-1 {
		s[i] , s[j] = s[j] , s[i]
	}
	//return s
}
func reversePointSlice(s *[]string) []string{
	i,j := 0 , len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j] , (*s)[i]
		i,j = i+1, j-1
	}
	return *s
}
func reversePointArray(s *[13]string) [13]string {
	n := len(*s)
	for i,j := 0 , n -1; i < j ; i,j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j] , (*s)[i]
	}
	return *s
}
func reversePointArray2(s [13]string) [13]string {
	n := len(s)
	for i,j := 0 , n -1; i < j ; i,j = i+1, j-1 {
		s[i], s[j] = s[j] , s[i]
	}
	return s
}
func main() {
	// slice传递 是引用传递, 所以操作会修改 底层数组, 导致源s会改变
	s := []string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	fmt.Println(s)		//[ Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec]

	reverse(s)
	fmt.Println(s)		//[Dec Nov Oct Sep Aug Jul Jun May Apr Mar Feb Jan ]

	reversePointSlice(&s)
	fmt.Println(s)		//[ Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec]

	// 指针方式数组传递 也会 改变源数组
	s2 := [13]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	fmt.Println(reversePointArray(&s2))	//[Dec Nov Oct Sep Aug Jul Jun May Apr Mar Feb Jan ]
	fmt.Println(s2)						//[Dec Nov Oct Sep Aug Jul Jun May Apr Mar Feb Jan ]

	s3 := [13]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	fmt.Println(reversePointArray2(s3))	// 已经反转
	fmt.Println(s3)						// 源数组未改变
	////var s1 []int
	////s1 := make([]int,1,10)[1:]
	//s1 := make([]int,0,10)
	////s1 := []int{}
	//if s1 == nil { fmt.Println( "空?",s1)}else {fmt.Println("非空?",s1)}
	//if len(s1) == 0 { fmt.Println( "空?",s1)}else {fmt.Println("非空?",s1)}
	//
	//
	//a := []int{1,2,3}
	//b := a
	//fmt.Printf("%v:\t%p,%v:\t%p\n",a,a,b,b)        //[1 2 3]:        0xc000090020,[1 2 3]:   0xc000090020
	//b[2] = 4
	//fmt.Printf("%v:\t%p,%v:\t%p\n",a,a,b,b)        //[1 2 4]:        0xc000090020,[1 2 4]:   0xc000090020
	//c := a[1:]
	//fmt.Printf("%v:\t%p,%v:\t%p\n",a,a,c,c)        //[1 2 4]:        0xc000090020,[2 4]:     0xc000090028
}
