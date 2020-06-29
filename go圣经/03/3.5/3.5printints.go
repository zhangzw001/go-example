package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i , v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		if _,err := fmt.Fprintf(&buf, "%d", v) ; err != nil { fmt.Println(err)}
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	a :=[]int{1,2,3}
	fmt.Println(a,intsToString(a))
	//fmt.Println(reflect.TypeOf(a),reflect.TypeOf(intsToString(a)))
}


