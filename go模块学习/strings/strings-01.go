package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	s := "   abcd222xxx dfadf hello world! good bye !   "
	// 1 Contains 字符串是否包含 指定字符串
	// func Contains(s string, substr string) bool
	fmt.Println(strings.Contains(s, "hello"))

	// 2 Join 字符串 + 操作, 性能更好
	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{s, s}, ""))

	// 3 index 找到字符串s中substr的位置
	// func Index(s string, substr string) int
	fmt.Println(strings.Index(s, "good bye"))

	// 4 repeat 重复打印
	//func Repeat(s string, count int) string
	fmt.Println(strings.Repeat(s, 2))

	// 5. replace 替换字符, 小于0 表示全部
	fmt.Println(strings.Replace(s, "  ", "", -1))

	// 6. split 分割成slice
	fmt.Println(strings.Split(s, "good")[0])

	// 7. Trim 去掉开头和结尾的字符
	fmt.Println(strings.Trim(s, "  "))

	// 8. Fields 去掉空格 并按照空格分割
	fmt.Println(strings.Fields(s))

	// 练习2：从Email中提取出用户名和域名：abc@163.com
	mail := "abc@163.com"
	m := strings.Split(strings.Replace(mail, " ", "", -1), "@")
	fmt.Printf("用户名是: %v, 域名是: %v\n", m[0], m[1])

	// 练习3: 让用户输入一句话,判断这句话中有没有邪恶,如果有邪恶就替换成这种形式然后输出,如:老王很邪恶,输出后变成老王很**
	//var str string
	//fmt.Scan(&str)
	//if strings.Contains(str,"邪恶") {
	//	str = strings.Replace(str,"邪恶","**",-1)
	//}
	//fmt.Println(str)
	//  练习4：把{“诸葛亮”,”鸟叔”,”卡卡西”,”卡哇伊”}变成诸葛亮|鸟叔|卡卡西|卡哇伊,然后再把|切割掉
	s4 := []string{"诸葛亮", "鸟叔", "卡卡西", "卡哇伊"}
	fmt.Println(strings.Join(s4, "|"))
	//练习5：让用户输入一句话,找出所有e的位置
	//var s5 string
	//fmt.Scan(&s5)
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		s5 := buf.Text()
		for i, v := range s5 {
			if v == 101 {
				fmt.Printf("%v ", i)
			}
		}
	}


    //练习6: 文本文件中存储了多个文章标题、作者，标题和作者之间用若干空格（数量不定）隔开，每行一个，标题有的长有的短，输出到控制台的时候最多标题长度10，如果超过10，则截取长度8的子串并且最后添加“...”，加一个竖线后输出作者的名字.

}
