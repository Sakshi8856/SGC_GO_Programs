package main

import "fmt"

func main() {
	// Example of a simple for loop
	fmt.Println("Using a basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example of an infinite loop (break when a condition is met)
	fmt.Println("\nUsing a loop with break:")
	count := 1
	for {
		if count > 3 {
			break
		}
		fmt.Println(count)
		count++
	}

	// Example of a for loop with continue (skip even numbers)
	fmt.Println("\nUsing continue to skip even numbers:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i)
	}
}
