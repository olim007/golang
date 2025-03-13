package main

import "fmt"

type node struct {
	value int
	next  *node
}

func printValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printValue(n.next)
	}
}

func mainPointer() {
	first := node{value: 0}
	second := node{value: 1}
	third := node{value: 2}
	first.next = &second
	second.next = &third

	printValue(&first)

	var current = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
