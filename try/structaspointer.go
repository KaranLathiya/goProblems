package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func modifyValue(point Point) Point {
	point.x += 10
	return point
}

func modifyPointer(point *Point) {
	point.x = 5
	point.y = 5
}

func modifyReference(point *Point) {
	point = &Point{5, 5}
}

func main() {
	p := Point{0, 0}
	fmt.Println(p) // prints (0, 0)

	p1 := modifyValue(p)
	fmt.Println(p1) // prints (0, 0)
	fmt.Println(p)
	modifyPointer(&p)
	fmt.Println(p) // prints (5, 5)

	p = Point{0, 0}
	modifyReference(&p)
	fmt.Println(p) // prints (0, 0)
}
