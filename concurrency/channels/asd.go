package main

import "fmt"

func main11() {
	ch := make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		ch <- 6
		<- ch
		fmt.Println("Go routine waiting receiver")
		//fmt.Scanln()
	}()
	ch <- 1
	go func() {
		fmt.Println("Go routine starts")
		num := <- ch
		ch <- 7
		fmt.Println("Go routine waiting receiver", num)
		//fmt.Scanln()
	}()
	//fmt.Scanln()
	go func() {
		fmt.Println("Go routine starts")
		ch <- 8
		fmt.Println("Go routine waiting receiver")
		//fmt.Scanln()
	}()
	fmt.Println(<- ch)
	//fmt.Scanln()

	fmt.Println("The end")
}
