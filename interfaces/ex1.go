package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	name string
	model string
}

type Aircraft struct {
	name string
	model string
}

func (c Car) move() {
	fmt.Println("Автомобил едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main1() {
	var car Vehicle = Car{}
	var plane Vehicle = Aircraft{}
	car.move()
	plane.move()

	bmv := Car{}
	bmv.model = "sport"
	bmv.name = "bmv"
	bmv.move()
}