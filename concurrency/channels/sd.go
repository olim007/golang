package main

import (
	"fmt"
)

func wdmain() {
	ch := make(chan int)

	go square(ch)
	ch <- 3
	fmt.Println("res:=", <-ch)
	fmt.Println("The end")
}

func square(ch chan int) {
	num := <- ch
	fmt.Println("num:=", num)
	ch <- num * num
}