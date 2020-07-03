package main

import (
	"fmt"
	"log"
)

//练习5.16：编写多参数版本的strings.Join。
func main() {
	//fmt.Println(strings.Join([]string{"123","abc"}," "))
	//fmt.Println(JoinNew(" ","A","b","C"))
	if _,err := JoinNew(","); err != nil {
		log.Fatalf("func JoinNew is err : %s",err)
	}
}

func JoinNew(sep string , a ...string )  (string , error)  {
	if len(a) == 0 {
		//log.Fatal("is nil")
		return "",fmt.Errorf(" is nil ")
	}
	var s string
	for _, i := range  a {
		if len(s) == 0 {
			s = i
			continue
		}
		s += fmt.Sprintf("%s%s",sep,i)
	}
	return s,nil

}
