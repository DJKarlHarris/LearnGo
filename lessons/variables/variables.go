package variables

import "fmt"

func Lesson() {
	fmt.Println("=== Variables Lesson ===")
	var name string = "Go Language"
	age := 25
	isStudent := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
}
