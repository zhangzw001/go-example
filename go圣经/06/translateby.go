package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1,2}
	q := Point{2,4}
	fmt.Println(p.Add(q))
	fmt.Println(p.Sub(q))
	path1 := Path{q}
	path1.TranslateBy(p,false)
	fmt.Println(path1)

	path1.TranslateBy(p,true)
	fmt.Println(path1)

	path2 := Path.TranslateBy
	path22 := Path{p}
	path2(path22,q,false)
	fmt.Println(path2)
}