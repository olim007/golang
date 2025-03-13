package main

import (
	"fmt"
	. "tutorials/models"
)

type library []string

func (l library) print() {
	for _, v := range l {
		fmt.Println(v, len(l))
	}
}

func main() {
	//short_struct()
	//fmt.Println()
	//long_struct()
	//mainPointer()

	var l = library{"B1", "B2", "B3"}
	l.print()

	var tom = Person{Name: "Tom", Age: 22}
	var kate = Person{
		Name: "Kate",
		Age:  21,
	}
	tom.Print()
	tom.Eat("stake")
	tom.UpdateAge(33)
	fmt.Println("tom", tom.Age)
	fmt.Println("kate", kate.Age)
}
