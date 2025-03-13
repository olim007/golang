package main

import "fmt"

type Form interface {
	perimeter() int
}

type Square struct {
	side int
}

func (s *Square) perimeter() int {
	return s.side * 4
}

type DecoratedForm struct {
	form       Form
	decoration string
}

func (d *DecoratedForm) perimeter() int {
	return d.form.perimeter()
}

func (d *DecoratedForm) getDecoration() string {
	return d.decoration
}

type TaggedForm struct {
	form Form
	tag  string
}

func (t *TaggedForm) perimeter() int {
	return t.form.perimeter()
}

func (t *TaggedForm) getTag() string {
	return t.tag
}

func maind() {
	s := &Square{5}
	d := &DecoratedForm{s, "fancy"}
	t := TaggedForm{d, "v1"}
	fmt.Println("Периметр:", t.perimeter())
	fmt.Println("Украшение:", t.form.(*DecoratedForm).getDecoration())
	fmt.Println("Тег:", t.getTag())
}
