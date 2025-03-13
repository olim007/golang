package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(10000)
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()
	time.Sleep(10000)
	fmt.Println("Hello from main")
}