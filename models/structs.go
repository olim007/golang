package models

import "fmt"

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Name string
	Age  int
	Contact
}

func (P Person) Print() {
	fmt.Println("Name", P.Name)
	fmt.Println("Age", P.Age)
}

func (P Person) Eat(Meal string) {
	fmt.Println(P.Name, "eat", Meal)
}

func (p *Person) UpdateAge(newAge int) {
	(*p).Age = newAge
}

type Person_ struct {
	Name        string
	Age         int
	ContactInfo Contact
}
