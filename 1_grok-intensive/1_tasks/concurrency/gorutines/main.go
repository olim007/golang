package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func artf(ch chan int) {
	for i := 0; i < 1000; i++ {
		if i == 678 {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go artf(ch)
	fmt.Println("wait 1 seconds")
	time.Sleep(time.Second)
	fmt.Println("wait 2 seconds")
	time.Sleep(time.Second)
	fmt.Println("wait 3 seconds")
	time.Sleep(time.Second)
	fmt.Println(<-ch)

	//go sayHello("ASF")
	//go sayHello("SDS")
	//time.Sleep(time.Second)
	//fmt.Println("Done")
	//x := 10
	//ch := make(chan int)
	//go func () {
	//	x = 20
	//	ch <- 42
	//}()
	//v := <- ch
	////b := <- ch
	//fmt.Println("v:", v)
	//fmt.Println("x:", x)
	//fmt.Println()
	//ch = make(chan int, 10)
	//ch1 := make(chan int, 2)
	//for i := 1; i <= 5; i++ {
	//	ch <- i
	//	//<- ch
	//}
	//fmt.Println(<- ch1)
	//
	//for i := 1; i <= 6; i++ {
	//	//ch <- i
	//	fmt.Println(<-ch)
	//}
}