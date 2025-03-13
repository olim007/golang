package __pointers

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main4() {
	a, b := 10, 20
	fmt.Println("До: ", a, b)
	swap(&a, &b)
	fmt.Println("После: ", a, b)
}
