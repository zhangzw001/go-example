package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {
	f ,err := os.Open("/tmp/test1.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
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


