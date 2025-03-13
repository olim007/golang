package main

import "fmt"

type Figure interface {
	area() int
}

type Box struct {
	width int
	height int
}

func (b *Box) area() int {
	return b.width * b.height
}

type PaintedFigure struct {
	figure Figure
	color  string
}

func (p *PaintedFigure) area() int {
	return p.figure.area()
}

func (p *PaintedFigure) getColor() string {
	return p.color
}

func main9() {
	b := &Box{3,4}
	p := PaintedFigure{b, "red"}
	fmt.Println("Площадь", p.area())
	fmt.Println("Цвет", p.getColor())
}