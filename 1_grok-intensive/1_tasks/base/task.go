package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) makeOlder() {
	(*p).age += 1
}

func makeOlder(p *Person) {
	p.age += 1
}

func main5() {
	alice := Person{"Alice", 25}
	fmt.Println("До: ", alice.age)
	alice.makeOlder()
	fmt.Println("После: ", alice.age)

	fmt.Println()

	fmt.Println("До: ", alice.age)
	makeOlder(&alice)
	fmt.Println("После: ", alice.age)

}
