package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Nartin Solveig", "Smash", 2011, length("4m24s")},
	{"Ready 3 Go", "Martin Solveig", "Smash", 2011, length("4m14s")},
}

type tablesSort struct {
	t 			[]*Track
	less 		func(i,j *Track ) bool
}
func (x tablesSort) Len() int			{ return len(x.t)}
func (x tablesSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x tablesSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
var clickRecords []string

func less(x,y *Track ) bool {
	for _, click := range clickRecords {
		if click == "Title" {
			if x.Title == y.Title {
				continue
			}
			return x.Title < y.Title
		}
		if click == "Year" {
			if x.Year == y.Year {
				continue
			}
			return x.Year < y.Year
		}
		if click == "Artist" {
			if x.Artist == y.Artist {
				continue
			}
			return x.Artist < y.Artist
		}
		if click == "Album" {
			if x.Album == y.Album {
				continue
			}
			return x.Album < y.Album
		}
		if click == "Length" {
			if x.Length == y.Length {
				continue
			}
			return x.Length < y.Length
		}
	}
	return false
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	//知识点
	//1.sort包内置的提供了根据一些排序函数来对任何序列排序的功能
	//2.Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设
	//3.Go使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定
	//4.一个内置的排序算法需要三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式
	//5.接口值得动态类型如果是可以比较的，即可以作为map的key或者switch的语句操作数
	//一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的
	//练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
	//定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。
	// 这种排序会保持第一个选择的列排完序的结果, 继续拍第二个选择的列(所以你会看到, 后面选择的列都在前面的结果之下)
	fmt.Println("------------------------------------------------")
	clickRecords = append(clickRecords, "Year")
	fmt.Println(clickRecords)
	sort.Sort(tablesSort{tracks,less})
	printTracks(tracks)
	fmt.Println("------------------------------------------------")
	clickRecords = append(clickRecords, "Artist")
	fmt.Println(clickRecords)
	sort.Sort(tablesSort{tracks,less})
	printTracks(tracks)
	fmt.Println("------------------------------------------------")
	clickRecords = append(clickRecords, "Title")
	fmt.Println(clickRecords)
	sort.Sort(tablesSort{tracks,less})
	printTracks(tracks)



}