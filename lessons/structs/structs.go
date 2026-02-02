package structs

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

// Lesson demonstrates Go structs
func Lesson() {
	fmt.Println("=== Structs Lesson ===")

	p1 := Person{Name: "John", Age: 30, City: "New York"}
	fmt.Printf("Person 1: %+v\n", p1)

	p2 := Person{"Jane", 25, "San Francisco"}
	fmt.Printf("Person 2: %+v\n", p2)
}
