package __pointers

import "fmt"

func double(x *int){
	*x *= 2
}

func main2() {
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}
