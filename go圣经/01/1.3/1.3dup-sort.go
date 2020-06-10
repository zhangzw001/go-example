package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}


func main() {
	start := time.Now()
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	fmt.Println("去重,未排序:")
	for line, n := range counts {
		fmt.Printf("%d\t\t%s\n",n, line)
	}

	var p1 PairList
	p1 = sortMapByValue(counts)
	//fmt.Println(p1)
	n := len(p1)
	fmt.Println("从大到小排序:")
	for i,_ := range p1 {
		j := p1[n-i-1]
		fmt.Printf("%d\t\t%s\n",j.Value,j.Key)
		//fmt.Println(i,n)
		//fmt.Println(p1[n-i-1])
	}
	fmt.Println("从小到大排序:")
	for _, j := range p1 {
		fmt.Printf("%d\t\t%s\n",j.Value,j.Key)
	}

	fmt.Printf("%.6f\n",time.Since(start).Seconds())
}
