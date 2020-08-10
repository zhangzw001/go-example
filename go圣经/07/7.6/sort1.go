package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Interface interface {
	Len()		int
	Less(i, j int) bool 	//索引
	Swap(i, j int )
}


type StringSlice []string
func (p StringSlice) Len() int 				{ return len(p) }
func (p StringSlice) Less(i,j int ) bool 	{ return p[i] < p[j]}
func (p StringSlice) Swap(i,j int )			{ p[i], p[j] = p[j], p[i]}



func main() {
	names := []string{"1","4","2"}
	sort.Sort(StringSlice(names))
	fmt.Println(reflect.TypeOf(names))
}