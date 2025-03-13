package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	area() int
	scale(factor int)
}

type Rectangle struct {
	width int
	height int
}

type Circle struct {
	radius int
}

func (r *Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func (c *Circle) area() int {
	return c.radius * c.radius * 3
}

func (c *Circle) scale(factor int) {
	c.radius *= factor
}

func createShape(kind string, a, b int) (Shape, error) {
	if kind == "rectangle" {
		return &Rectangle{a, b}, nil
	} else if kind == "circle" {
		return &Circle{a}, nil
	}
	return nil, errors.New("unknown shape")
}

func processShape(s Shape, err error, name string) {
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	s.scale(2)
	fmt.Println("Площадь", name, s.area())
}

func mainh() {
	r, errR := createShape("rectangle", 2, 3)
	processShape(r, errR, "прямоугольника")
	c, errC := createShape("circle", 2, 3)
	processShape(c, errC, "круга")
	t, errT := createShape("triangle", 2, 3)
	processShape(t, errT, "треугольника")
}

