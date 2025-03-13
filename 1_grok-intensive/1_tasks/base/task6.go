package main

import "fmt"

//func (r *Rectangle) scale(factor int) {
//	r.width *= factor
//	r.height *= factor
//}

func maink() {
	r := Rectangle{4,2}
	fmt.Println("До:", r.width, r.height)
	r.scale(3)
	fmt.Println("После:", r.width, r.height)
}
