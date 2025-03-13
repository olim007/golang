package main

import "fmt"

func allocateNum(x **int) {
	*x = new(int)
	**x = 42
}

func main8() {
	var num *int = nil
	fmt.Println("До: ", num)
	allocateNum(&num)
	fmt.Println("После: ", *num)
}
