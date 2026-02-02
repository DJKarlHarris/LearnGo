package main

import (
	"fmt"
	"learn-go/lessons/arrays"
	"learn-go/lessons/channels"
	"learn-go/lessons/functions"
	"learn-go/lessons/goroutines"
	"learn-go/lessons/interfaces"
	"learn-go/lessons/maps"
	"learn-go/lessons/slices"
	"learn-go/lessons/structs"
	"learn-go/lessons/variables"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <lesson-name>")
		fmt.Println("Available lessons:")
		fmt.Println("  - variables")
		fmt.Println("  - arrays")
		fmt.Println("  - slices")
		fmt.Println("  - maps")
		fmt.Println("  - structs")
		fmt.Println("  - functions")
		fmt.Println("  - interfaces")
		fmt.Println("  - goroutines")
		fmt.Println("  - channels")
		return
	}

	lesson := os.Args[1]

	switch lesson {
	case "variables":
		variables.Lesson()
	case "arrays":
		arrays.Lesson()
	case "slices":
		slices.Lesson()
	case "maps":
		maps.Lesson()
	case "structs":
		structs.Lesson()
	case "functions":
		functions.Lesson()
	case "interfaces":
		interfaces.Lesson()
	case "goroutines":
		goroutines.Lesson()
	case "channels":
		channels.Lesson()
	default:
		fmt.Printf("Unknown lesson: %s\n", lesson)
		fmt.Println("Available lessons:")
		fmt.Println("  - variables")
		fmt.Println("  - arrays")
		fmt.Println("  - slices")
		fmt.Println("  - maps")
		fmt.Println("  - structs")
		fmt.Println("  - functions")
		fmt.Println("  - interfaces")
		fmt.Println("  - goroutines")
		fmt.Println("  - channels")
	}
}
