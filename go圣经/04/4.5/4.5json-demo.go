package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Demo struct {
	Id 		int64 `json:"id"`
	Name 	string `json:"name"`
}

type Demo2 struct {
	Num 	int `json:"num"`
	Title   string `json:"title"`
	Img 	string `json:"img"`
}
func main() {
	//区别
	//1、json.NewDecoder是从一个流里面直接进行解码，代码精干
	//2、json.Unmarshal是从已存在与内存中的json进行解码
	//3、相对于解码，json.NewEncoder进行大JSON的编码比json.marshal性能高，因为内部使用pool
	//
	//场景应用
	//1、json.NewDecoder用于http连接与socket连接的读取与写入，或者文件读取
	//2、json.Unmarshal用于直接是byte的输入
	fmt.Println("---------------------------------------------------")
	demojson1 := `{
	"id":123,
	"name":"demo",
	"other":"none"
	}`

	// NewDecoder用法
	var demo Demo
	json.NewDecoder(strings.NewReader(demojson1)).Decode(&demo)
	fmt.Println(demo.Name)

	// MarshalIndent 或Marshal
	data, err := json.MarshalIndent(demo,"", " ")
	if err != nil { fmt.Println(err)}
	fmt.Println(reflect.TypeOf(data),len(data),data)
	fmt.Printf("%s\n", data)


	fmt.Println("---------------------------------------------------")

	// http.get
	resp , err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil { fmt.Println( err )}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("http.get error : %d\n",resp.Status)
	}
	var demo2 Demo2
	if err := json.NewDecoder(resp.Body).Decode(&demo2); err != nil {
		resp.Body.Close()
		fmt.Println( err )
	}

	fmt.Println(demo2.Img)

}
