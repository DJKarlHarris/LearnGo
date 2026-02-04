package slices

import "fmt"

func Lesson() {
	fmt.Println("=== Slices Lesson ===")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Length: %d\n", len(slice))
	fmt.Printf("Capacity: %d\n", cap(slice))

	slice = append(slice, 6)
	fmt.Printf("After append: %v\n", slice)

	slen := len(slice)
	fmt.Printf("slen: %d Slice: %v\n", slen, slice[:slen])
	fmt.Printf("slen: %d Slice: %v\n", slen, slice[slen-1:])
}
