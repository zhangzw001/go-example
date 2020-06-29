package main

import (
	github "github.com/zhangzw001/go-example/go圣经/04/4.5/json-github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
// |操作符表示将前一个表达式的结果作为后一个函数的输入, .CreatedAt作为参数传递到 daysAgo函数

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24 )
}

func main() {
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo":daysAgo}).
		Parse(templ))
	//report , err := template.New("report").
	//	Funcs(template.FuncMap{"daysAgo":daysAgo}).
	//	Parse(templ)
	//if err != nil { log.Fatal(err)}

	result , err := github.SearchIssues(os.Args[1:])
	if err != nil { log.Fatal(err)}
	if err := report.Execute(os.Stdout, result ); err != nil {
		log.Fatal(err )
	}
}
