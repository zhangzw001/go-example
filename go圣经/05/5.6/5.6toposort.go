package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	//for i, course := range topSort(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}


	topoSort := topSort2(prereqs)
	var s = make([]int,0)
	for k,_ := range topoSort {
		s = append(s, k)
	}
	sort.Ints(s)
	for i , _ := range s {
		fmt.Printf("%d:\t%s\n",i+1, topoSort[i])
	}


}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topSort2(m map[string][]string) map[int]string {
	var order = make(map[int]string)

	seen := make(map[string]bool)

	var visitAll2 func(items []string)
	index := 0
	visitAll2 = func( items []string ) {
		for _,item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll2(m[item])
				order[index] = item
				index++
			}
		}
	}

	var keys  []string
	for k,_ := range m {
		keys = append(keys,k)
	}
	//sort.Strings(keys)
	visitAll2(keys)

	return order
}

func topSort3(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	index := 0
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}