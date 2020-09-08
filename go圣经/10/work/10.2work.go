package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//练习 10.2： 设计一个通用的压缩文件读取框架，用来读取ZIP（archive/zip）和POSIX tar（archive/tar）格式压缩的文档。
//使用类似上面的注册技术来扩展支持不同的压缩格式，然后根据需要通过匿名导入选择导入要支持的压缩格式的驱动包。

func main() {
	filename := "/tmp/a.txt.zip"
	f,_ := os.Open(filename)
	defer f.Close()
	fmt.Println(filepath.Ext(filename))
	abs,_ := filepath.Abs("a.txt")
	fmt.Println(abs)
	fmt.Println(filepath.Dir(filename))


}



func toDecompressor(zipfile string) {

	read, err := zip.OpenReader(zipfile)
	defer read.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range read.File {
		//如果是目录则创建
		//if f.FileInfo().IsDir() {
		//	if err = os.Mkdir(f.Name,f.Mode()); err != nil {
		//		log.Fatal(err)
		//	}
		//	continue
		//}

		//如果是文件
		fmt.Println(f.Name)
	}
}