package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	key , _ := MD5("abcd"+time.Now().String())
	fmt.Println(key)
	key , _ = MD5("abcd")
	fmt.Println(key)
}


func MD5(s string) (string ,error) {
	secret := md5.New()
	_,err := secret.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(secret.Sum(nil)),nil
}
