package channels

import "fmt"

func Lesson() {
	fmt.Println("=== Channels Lesson ===")
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	fmt.Printf("Channel content: %s, %s\n", <-ch, <-ch)
}
