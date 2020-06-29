package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func flagSha256() {
	flag1 	:= 	flag.String("f","SHA256", "select hash method(SHA256,SHA384,SHA512)")
	s 		:= 	flag.String("s","","a string")
	flag.Parse()
	switch *flag1 {
	case "SHA384":
		fmt.Printf("%x\n",sha512.Sum384([]byte(*s)))
	case "SHA512":
		fmt.Printf("%x\n",sha512.Sum512([]byte(*s)))
	case "SHA256":
		fmt.Printf("%x\n",sha256.Sum256([]byte(*s)))
	default:
		fmt.Println("no support method")
	}

}

func main() {
	flagSha256()
}