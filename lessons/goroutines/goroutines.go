package goroutines

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond) // Adding delay to see the effect
	}
}

func printLetters() {
	for char := 'a'; char <= 'c'; char++ {
		fmt.Printf("Letter: %c\n", char)
		time.Sleep(100 * time.Millisecond) // Adding delay to see the effect
	}
}

func Lesson() {
	fmt.Println("=== Goroutines Lesson ===")
	fmt.Println("Main function starts")

	go printNumbers()
	go printLetters()

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Main function ends")
}
