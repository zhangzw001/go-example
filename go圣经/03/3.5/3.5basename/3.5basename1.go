package basename

import (
	"strings"
)

func Basename1(s string ) string {
	for i:=len(s)-1; i >=0; i -- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s)-1 ; i >= 0 ; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func Basename2(s string ) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
//func main() {
//	s1 := "/data/gatlingCA.key.pem"
//	fmt.Println(basename1(s1))
//	fmt.Println(basename2(s1))
//}