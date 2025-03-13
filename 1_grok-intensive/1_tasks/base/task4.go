package main

import "fmt"

func mains() {
	a := 1
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println()
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)

}
