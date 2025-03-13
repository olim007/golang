package main

import "fmt"

type Outline interface {
	length() int
}

type Frame struct {
	width int
	height int
}

type StyleOutline struct {
	outline Outline
	style   string
}

func (styleOutline *StyleOutline) length() int {
	return styleOutline.outline.length()
}

func (styleOutline *StyleOutline) getStyle() string {
	return styleOutline.style
}

func createOutline(kind string, a, b int) Outline {
	if kind == "frame" {
		return &Frame{a, b}
	}
	return nil
}

func (f *Frame) length() int {
	return 2 * (f.width + f.height)
}

func mainqw() {
	f := createOutline("frame", 3, 4)
	styleOutline := StyleOutline{f, "dash"}
	fmt.Println("Длина контура:", styleOutline.length())
	fmt.Println("Стиль:", styleOutline.getStyle())
}
