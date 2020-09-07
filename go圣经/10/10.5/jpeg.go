package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {
	f ,err := os.Open("test1.png")
	if err != nil {
		log.Fatal("文件读取失败")
	}
	defer f.Close()
	fmt.Println(f.Name())
	if err := toJPEG(f,os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr,"jpeg: %v\n", err)
		os.Exit(1)
	}

}


func toJPEG(in io.Reader , out io.Writer) error {
	img, kind , err := image.Decode(in)
	if err != nil { return err}
	fmt.Fprintf(os.Stderr, "Input format=%v\n",kind)
	return jpeg.Encode(out,img,&jpeg.Options{Quality:95})

}


