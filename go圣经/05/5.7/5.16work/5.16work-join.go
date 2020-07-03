package main

import (
	"fmt"
	"strings"
)

//练习5.16：编写多参数版本的strings.Join。
func main() {
	fmt.Println(strings.Join([]string{"123","abc"}," "))
}
