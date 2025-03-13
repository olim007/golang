package main

import "fmt"

type Contour interface {
	size() int
}

type Panel struct {
	length int
	width int
}

func (p *Panel) size() int {
	return 2 * (p.width + p.length)
}

type TexturedContour struct {
	contour Contour
	texture string
}

func (t *TexturedContour) size() int {
	return t.contour.size()
}

func (t *TexturedContour) getTexture() string {
	return t.texture
}

func createContour(kind string, a, b int) Contour {
	if kind == "panel" {
		return &Panel{a, b}
	}
	return nil
}

func main() {
	p := createContour("panel", 5, 2)
	t := TexturedContour{p, "wood"}
	var c Contour = &t

	fmt.Println("Размер:", c.size())
	if textured, ok := c.(*TexturedContour); ok {
		fmt.Println("Текстура:", textured.getTexture())
	} else {
		fmt.Println("Текстура недоступна")
	}
}
