package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。

	fmt.Println(strings.Join(os.Args[:]," "))
	//练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

	for k,v := range os.Args {
		fmt.Printf("索引: %v\t值: %v\n",k,v)
	}
	//练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
	start := time.Now()
	for i := 0 ; i < 200000 ; i ++ {
		test_for(os.Args)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())

	start = time.Now()
	for i := 0 ; i < 200000 ; i ++ {
		strings.Join(os.Args[:], " ")
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())


}


func test_for(slice []string ) string {
	var s1, sep1 string
	for i :=0 ;i < len(slice) ; i ++ {
		s1 += sep1 + slice[i]
		sep1 = " "
	}
	return s1
}


