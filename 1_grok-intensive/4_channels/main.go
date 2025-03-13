package main

import "fmt"

func main() {
	r := make([]int, 2)
	ch := make(chan *int)
	fmt.Println(r)
	fmt.Println(ch)
}