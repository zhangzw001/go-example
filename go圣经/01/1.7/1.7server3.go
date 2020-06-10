package main

import (
	"github.com/zhangzw001/go-example/go圣经/01/1.4"
	"image/color"
	"log"
	"net/http"
)


var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
func main() {

	http.HandleFunc("/gif",func(w http.ResponseWriter,r *http.Request){ 	lissajous.Lissajous(w)})
	log.Fatal(http.ListenAndServe("localhost:8000",nil ))
}


