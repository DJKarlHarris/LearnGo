package arrays

import "fmt"

func Lesson() {
	fmt.Println("=== Arrays Lesson ===")
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Length: %d\n", len(arr))

	// Array literal
	colors := [3]string{"red", "green", "blue"}
	fmt.Printf("Colors: %v\n", colors)
}
