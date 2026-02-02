package functions

import "fmt"

func Lesson() {
	fmt.Println("=== Functions Lesson ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	sum, product := calculate(4, 6)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}

func add(a, b int) int {
	return a + b
}

func calculate(x, y int) (int, int) {
	return x + y, x * y
}
