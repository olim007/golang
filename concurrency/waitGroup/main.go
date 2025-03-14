package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(3)
	work := func(id int) {
		fmt.Println(<-ch)
		defer wg.Done()
		fmt.Println("Горутина", id, "начала выполнение \n")
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина", id, "завершила выполнение \n")
	}

	go work(1)
	ch <- 1
	go work(2)
	ch <- 2
	go work(3)
	ch <- 3

	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
