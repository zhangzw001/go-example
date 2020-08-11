package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//练习 7.11： 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。
//例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
//（注意：这个修改会引入变量同时更新的问题）

//知识点
//1.单一使用handler处理
//2.聚合使用handlers处理
//3.http.HandlerFunc是一个类型，实现了接口http.Handler方法的函数类型
//4.net/http包提供了一个全局的ServeMux实例DefaultServerMux

func main() {
	//所以为了方便，net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主handler，我们不需要将它传给ListenAndServe函数；nil值就可以工作。
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/update",db.update)
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))
	//log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars


func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _,ok := db[item] ; !ok {
		fmt.Fprintf(w, "no such item : %v\n",item )
		return
	}
	fprice, err := strconv.ParseFloat(price,32)
	if err != nil  {
		fmt.Fprintf(w,"price is not float: %v\n",price)
		return
	}
	fmt.Fprintf(w,"item: %v  old price : %v\n",item, db[item])
	db[item] = dollars(fprice)
	fmt.Fprintf(w,"item: %v  new price : %v\n",item, db[item])


}
func (db database) list(w http.ResponseWriter, req *http.Request) {
	var shopList = template.Must(template.New("shopList").Parse(`
	<h1>shopList</h1>
	<table border="2">
	<tr style='text-align:left'>
	<th>item</th>
	<th>	</th>
	<th>price</th>
	</tr>
	</table>
`))
	shopList.Execute(w,nil)
	const templ = `<p>{{.A}}	{{.B}}</p>`
	type data struct {
		A 	string
		B 	dollars
	}

	t := template.Must(template.New("escape").Parse(templ))
	for item, price := range db {
		var dat = data{item,price}
		err := t.Execute(w, dat)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}