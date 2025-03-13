package main

import "fmt"

func send(ch chan <- int) {
	ch <- 22333
}

func _main() {
	ch := make(chan int)
	go send(ch)
	fmt.Println(<-ch)

	//var send_ch chan <- int
	////var get_ch <-chan int
	//
	//send_ch <- 2
	//send_ch <-2
	//ch := make(chan int, 3)
	//ch <- 2
	//fmt.Println(len(ch))
	//ch <- 22
	//fmt.Println(len(ch))
	//ch <- 12
	//fmt.Println(len(ch))
	//if len(ch) < 3 {
	//	ch <- 123
	//}
	//fmt.Println(len(ch))
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
