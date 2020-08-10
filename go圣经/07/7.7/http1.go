package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//http1
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
	//----------------------------------
	//http2
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w,"%s:\t%s\n",item,price)
		}
	case "/price":

		//fmt.Println(req.URL.Query())	//http://localhost:8000/price?item=shoes&item2=aa -> map[item:[shoes] item2:[aa]]
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"no such item:\t%q\n",item )
			return
		}
		fmt.Fprintf(w,"%s\n",price)
	default:
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w,"no such page:\t%s\n",req.URL)
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound) // 404
	}
}
