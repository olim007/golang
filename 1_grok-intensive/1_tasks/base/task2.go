package main

import "fmt"

func resetToZero(x **int) {
	**x = 0
}

func main7() {
	x := 7
	p := &x
	fmt.Println("До: ", x)
	resetToZero(&p)
	fmt.Println("После: ", x)
}
