package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title		string
	Year 		int `json:"released"`	//Year名字的成员在编码后变成了released,一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串：
	Color		bool `json:"color,omitempty"`		//Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。
	Actors 		[]string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}
	//这样的数据结构特别适合JSON格式，并且在两种之间相互转换也很容易。
	//将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组（marshaling）。
	//编组通过调用json.Marshal函数完成：
	data , err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}
	fmt.Printf("%s\n", data)
	//这种紧凑的表示形式虽然包含了全部的信息，但是很难阅读。
	//为了生成便于阅读的格式，另一个json.MarshalIndent函数将产生整齐缩进的输出。
	//该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：
	data2 , err := json.MarshalIndent(movies,""," ")
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}
	fmt.Printf("%s\n", data2)

	//在编码时，默认使用Go语言结构体的成员名字作为JSON的对象（通过reflect反射技术，我们将在12.6节讨论）。
	//只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字母开头的成员名称。

	//细心的读者可能已经注意到，其中Year名字的成员在编码后变成了released，还有Color成员编码后变成了小写字母开头的color。这是因为构体成员Tag所导致的。一个构体成员Tag是和在编译阶段关联到该成员的元信息字符串：

	//编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成。下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。通过定义合适的Go语言数据结构，我们可以选择性地解码JSON中感兴趣的成员。当Unmarshal函数调用返回，slice将被只含有Title信息值填充，其它JSON成员将被忽略。
	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}


