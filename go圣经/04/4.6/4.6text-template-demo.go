package main

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"
)

type demo struct {
	Id 		int `json:"id"`
	Name 	string `json:"name"`
	Other   string `json:"other"`
}

const demotemplate = `demo :
---------------------------
Id: 	{{.Id}}
Name:	{{.Name}}
Other:	{{.Other}}
`

func main() {
	demojson1 := `{
	"id":123,
	"name":"demo",
	"other":"none"
	}`

	//首先对json进行解析
	var demoresult demo
	json.NewDecoder(strings.NewReader(demojson1)).Decode(&demoresult)

	//其次生成模板
	//template.New先创建并返回一个模板；Funcs方法将daysAgo等自定义函数注册到模板中，并返回模板；最后调用Parse函数分析模板。
	//template.Must辅助函数可以简化这个致命错误的处理：它接受一个模板和一个error类型的参数，检测error是否为nil（如果不是nil则发出panic异常），然后返回传入的模板。
	report := template.Must(template.New("demo").
		Parse(demotemplate))

	report.Execute(os.Stdout, demoresult)
}
