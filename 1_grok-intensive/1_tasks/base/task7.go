package main

import "fmt"

//type Shape interface {
//	area() int
//	scale(factor int)
//}

//type Rectangle struct {
//	width int
//	height int
//}

//type Circle struct {
//	radius int
//}

//func (c *Circle) area() int {
//	return c.radius * c.radius * 3
//}

//func (c *Circle) scale(factor int) {
//	c.radius *= factor
//}

//func (r *Rectangle) area() int {
//	return r.width * r.height
//}

//func (r *Rectangle) scale(factor int) {
//	r.width *= factor
//	r.height *= factor
//}

func createRectangle(kind string, a, b int) Shape {
	if kind == "rectangle" {
		return &Rectangle{a, b}
	} else if kind == "circle" {
		return &Circle{a}
	} else {
		return nil
	}
}

func mainw() {
	//var sl []Shape = []Shape{Rectangle{4,2}, Circle{3}}
	//for i, shape := range sl {
	//	if i == 0 {
	//		fmt.Println("Площадь Rectangle:", shape.area())
	//	} else {
	//		fmt.Println("Площадь Circle:",shape.area())
	//	}
	//}
	//s := &Rectangle{2, 3}
	//s.scale(2)
	r := createRectangle("rectangle",2,3)
	r.scale(2)
	fmt.Println("Площадь прямоугольника:", r.area())
	c := createRectangle("circle",2,3)
	c.scale(2)
	fmt.Println("Площадь круга:", c.area())
}
