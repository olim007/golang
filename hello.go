package main

import (
	"fmt"
	"rsc.io/quote"
)

func mains() {
	message := quote.Hello()
	fmt.Println(message)
}
