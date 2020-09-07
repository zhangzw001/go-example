package work

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
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

	var type string
	type : = flag.String(type,"jpeg","输入图片格式")
	if err := toImg(f,os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr,"jpeg: %v\n", err)
		os.Exit(1)
	}


}


func toImg(in io.Reader , out io.Writer) error {
	img, kind , err := image.Decode(in)
	if err != nil { return err}
	fmt.Fprintf(os.Stderr, "Input format=%v\n",kind)

	return jpeg.Encode(out,img,&jpeg.Options{Quality:95})

}


