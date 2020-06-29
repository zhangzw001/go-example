package comma

import (
	"bytes"
	"strings"
)

func Comma(s string ) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func Comma2( s string ) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	for i :=3 ; i < n ; i +=3 {
		s = strings.Join([]string{s[:n-i],s[n-i:]},",")
	}
	return s
}

func Comma3( s string ) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	for i :=3 ; i < n ; i +=3 {
		s = s[:n-i] + "," + s[n-i:]
	}
	return s
}


func CommaBuffer(s string  ) string {
	var buf bytes.Buffer

	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	arr := strings.Split(s,".")
	s = arr[0]
	n := len(s)
	p := 3
	m := n%p
	if n <= p { return s}

	for i := 0 ; i < n ; i ++ {
		if (i-m)%p == 0 && i != 0 {
			//buf.WriteString(",")
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	if len(arr) >1 {
		buf.WriteByte('.')
		buf.WriteString(arr[1])
	}
	return buf.String()
}


