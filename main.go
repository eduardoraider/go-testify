package main

import "fmt"

// Function that adds two integers

func Sum(a, b int) int {
	return a + b
}

func main() {
	// Example usage of the Add function
	sum := Sum(3, 4)
	fmt.Println("The sum is:", sum)
}
