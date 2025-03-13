package main

import (
	"fmt"
	. "tutorials/models"
)

func longStruct() {
	var tom = Person_{
		Name: "Tom",
		Age:  22,
		ContactInfo: Contact{
			Email: "tom@gmail.com",
			Phone: "+12456786542",
		},
	}
	fmt.Println(tom.ContactInfo.Email)
	fmt.Println(tom.ContactInfo.Phone)
	fmt.Println(tom.Name)
}
