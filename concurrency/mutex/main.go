package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	ch := make(chan bool)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The end")
}

func work(n int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for i := 1; i <= 5; i++ {
		counter++
		fmt.Println("Goroutine", n, "-", counter)
	}
	mutex.Unlock()
	ch <- true
}
