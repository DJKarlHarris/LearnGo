package interfaces

import "fmt"

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Type() string {
	return "Rectangle"
}

type Shapes []Shape

func (ss Shapes) Area() float64 {
	var area float64
	for _, s := range ss {
		area += s.Area()
	}

	return area
}

func (ss Shapes) Type() string {
	return "Shapes"
}

func Lesson() {
	fmt.Println("=== Interfaces Lesson ===")

	r1 := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Shape: %s, Area: %.2f\n", r1.Type(), r1.Area())

	r2 := Rectangle{Width: 5, Height: 5}
	fmt.Printf("Shape: %s, Area: %.2f\n", r2.Type(), r2.Area())

	ss := Shapes{r1, r2}
	fmt.Printf("Shape: %s, Area: %.2f\n", ss.Type(), ss.Area())
}
