package main

import (
	"fmt"
	"github.com/zhangzw001/go-example/go圣经/01/1.4"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)



func main() {

	rand.Seed(time.Now().UnixNano())
	lconf := lissajous.Lconfig {
		Cycles	: 5,
		Res 	: 0.001,
		Freq 	: rand.Float64() * 3.0,
		Size 	: 100,
		Nframes  : 64,
		Delay 	: 8,
	}


	http.HandleFunc("/gif",func(w http.ResponseWriter,r *http.Request){

		confs := r.URL.Query()
		fmt.Fprintf(os.Stdout,"%v\n",confs)
		for i, c := range confs {
			switch i {
			case "cycles" : lconf.Cycles, _ = strconv.ParseFloat( c[0], 64 )
			case "freq"   : lconf.Freq  , _ = strconv.ParseFloat( c[0], 64 )
			case "res"    : lconf.Res   , _ = strconv.ParseFloat( c[0], 64 )
			case "size"   : lconf.Size  , _ = strconv.Atoi( c[0] )
			case "frames" : lconf.Nframes, _ = strconv.Atoi( c[0] )
			case "delay"  : lconf.Delay , _ = strconv.Atoi( c[0] )
			}
		}

		lissajous.Lissajous(w,lconf)
	})
	log.Fatal(http.ListenAndServe("localhost:8000",nil ))
}


