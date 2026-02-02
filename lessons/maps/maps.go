package maps

import "fmt"

func Lesson() {
	fmt.Println("=== Maps Lesson ===")
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("Alice's age: %d\n", ages["Alice"])

	// Check if key exists
	if age, exists := ages["David"]; exists {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Println("David not found in map")
	}
}
