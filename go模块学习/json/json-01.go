package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type Student struct {
	Name 	string
	Age		int
}
func main() {

	//json1 := `{
  // "name": "4acfe603-bef4-418d-9254-80c24dc78f4b",
  //  "age": 1,
  //}`
  //
	//var tmp Student
	//err := json.Unmarshal([]byte(json1), &tmp)
  //
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
  //
	//fmt.Println(tmp.Name)

	//
	json2 := `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"]}`

	if !gjson.Valid(json2) {
		fmt.Println("invalid json")
	}

	fmt.Println(gjson.Get(json2,`children`).Array()[0])

	// json3
	json3 := `{"channel":"shop"}`
	fmt.Println(gjson.Get(json3,`channel`))

}
