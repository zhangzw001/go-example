package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func createFile(path string ) {
	// 创建文件, 存在则删除
	//f, err := os.Create(path)
	// 默认打开只读的
	//f, err := os.Open(path)
	// 打开文件,设置可写
	// O_RDONLY(只读模式)，O_WRONLY(只写模式), O_RDWR( 可读可写模式)，O_APPEND(追加模式), O_CREATE(创建)。
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	// 关闭文件
	defer f.Close()

	if err != nil { log.Fatal(err )}
	// 自己添加换行
	str := fmt.Sprintln("this is writerString test")
	// 1. writerString 写入
	_, err = f.WriteString(str)
	if err == nil {
		fmt.Println("写入成功")
	}else {
		log.Fatal(err )
	}
	//// 2. write 写入
	//buf := []byte(str)
	//_, err = f.Write(buf)

	//// 3. writerat 写入
	//// 查找文件末尾的偏移量
	//n, _ := f.Seek(0,os.O_APPEND)
	//// 从末尾开始写入
	//_, err = f.WriteAt(buf,n)
	//if err != nil { log.Fatal(err )}

}

func ReadFile(path string) {
	// 打开文件
	f, _ := os.Open(path)
	defer f.Close()
	//
	buf := make([]byte, 1024)
	// n代表从文件读取内容的长度
	n , err  := f.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err )
	}
	fmt.Printf("buf = %v\n", string(buf[:n]))
}
func main() {
	createFile("file1.txt")
	ReadFile("file1.txt")
}