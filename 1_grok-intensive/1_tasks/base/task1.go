package main

import "fmt"

type Point struct {
	x, y int
}

func newPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func move(point *Point) {
	point.y += 3
	point.x += 2
}

func main6() {
	point := newPoint(1,4)
	fmt.Println("До: ", point.x, point.y)
	move(point)
	fmt.Println("После: ", point.x, point.y)
}
