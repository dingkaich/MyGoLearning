package main

import (
	"fmt"
	"image/color"
	"myprint"
)

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

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
	fmt.Println("good")
	myprint.Myprint()
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
	p := Path{
		{1, 2},
		{2, 3},
	}
	q := Point{1, 1}
	m := make(Path, 5)
	m = Path{
		{1, 2},
		{3, 4},
	}

	p.TranslateBy(q, true)
	m.TranslateBy(q, true)
	for i := range m {
		fmt.Println(m[i])
	}
	for _, i := range p {
		fmt.Println(i)
	}

	myprint.TestGoRange()
	myprint.TestInterface()
	myprint.Mychan()
	myprint.Mychan2()

}
