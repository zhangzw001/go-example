package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

//练习 10.1： 扩展jpeg程序，以支持任意图像格式之间的相互转换，使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式。


func main() {
	f ,err := os.Open("/tmp/test1.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var typeImg = flag.String("type","jpeg","输入图片格式")
	if err := toImg(f,os.Stdout,*typeImg); err != nil {
		fmt.Fprintf(os.Stderr,"jpeg: %v\n", err)
		os.Exit(1)
	}


}


func toImg(in io.Reader , out io.Writer, typeImg string) error {
	img, kind , err := image.Decode(in)
	if err != nil { return err}
	fmt.Fprintf(os.Stderr, "Input format=%v\n",kind)
	switch typeImg {
	case "jpeg" :
		return jpeg.Encode(out,img,&jpeg.Options{Quality:95})
	case "png" :
		return png.Encode(out,img)
	case "gif" :
		return gif.Encode(out,img,&gif.Options{})
	default:
		log.Fatal("不支持的转换格式")
	}
	return nil

}


