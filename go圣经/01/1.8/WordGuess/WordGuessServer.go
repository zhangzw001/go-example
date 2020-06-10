package main

import (
	"fmt"
	WordGuess "github.com/zhangzw001/go-example/go圣经/01/1.8/WordGuess/src"
	"html/template"
	"log"
	"net/http"
	"strconv"
)
type CustomMux struct {

}
func main() {

	//http.HandleFunc("/guess",func(w http.ResponseWriter, r *http.Request) {
	//	WordGuess.WordGuess(w)
	//})
	//指定访问的路由
	http.HandleFunc("/index",indexHandler)
	http.HandleFunc("/guess",guessHandler)
	//设定监听端口
	log.Fatal(http.ListenAndServe("0.0.0.0:8000",nil))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("src/view/index.html")
		checkErr(err)
		checkErr(t.Execute(w, nil))

}
func guessHandler(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("src/view/guess.html")
		checkErr(err)
		if r.Method == "GET" {
			checkErr(t.Execute(w, nil))
		} else {
			checkErr(t.Execute(w, nil))
			guessValue := r.FormValue("guessValue")
			guessValueNew, err := strconv.Atoi(guessValue)
			checkErr(err)
			// 每次传递一个 web表单输入的值
			WordGuess.WordGuess(w, guessValueNew)
			if WordGuess.Flag == true {
				t2,err := template.ParseFiles("src/view/guessReturn.html")
				checkErr(err)
				checkErr(t2.Execute(w,nil))
				WordGuess.Flag = false
				fmt.Println(WordGuess.RandValue)
			}
		}
}