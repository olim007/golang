package __pointers

import "fmt"

func add5(x *int) {
	*x += 5
}

func main3() {
	a := 3
	fmt.Println("До: ", a)
	add5(&a)
	fmt.Println("После: ", a)
}
