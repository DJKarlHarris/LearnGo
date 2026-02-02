package interfaces

import "fmt"

type Shape interface {
	Area() float64
	Type() string
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Type returns the type of the shape
func (r Rectangle) Type() string {
	return "Rectangle"
}

// Lesson demonstrates Go interfaces
func Lesson() {
	fmt.Println("=== Interfaces Lesson ===")

	r := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Shape: %s, Area: %.2f\n", r.Type(), r.Area())
}
