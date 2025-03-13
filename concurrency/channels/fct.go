package main

import "fmt"

//func factorial(n int, ch chan struct{}, results map[int]int) {
//	defer close(ch)
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//		results[i] = result
//	}
//}



func main() {

	//ch := make(chan struct{})
	//res := make(map[int]int)
	//
	//go factorial(5, ch, res)
	//
	//<-ch
	//
	//for i,v := range res {
	//	fmt.Println(i, " - ", v)
	//}

	ch := make(chan int)

	go factorial(7, ch)

	for {
		num, opened := <- ch
		if !opened {
			break
		}
		fmt.Println(num)
	}

	//fmt.Println(<-ch)
}

func factorial(n int, ch chan int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		ch <- result
	}
	//ch <- result
}
