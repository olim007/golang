package main

import (
	"fmt"
	. "tutorials/models"
)

func short_struct() {
	var tom = Person{
		Name: "Tom",
		Age:  22,
		Contact: Contact{
			Email: "tom@mail.ru",
			Phone: "+1023493212",
		},
	}
	fmt.Println(tom.Name)
	fmt.Println(tom.Age)
	fmt.Println(tom.Phone)
	fmt.Println(tom.Email)
}
