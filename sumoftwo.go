package main

import "fmt"

func main() {
	var a, b int

	// Input values
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	// Calculate sum
	sum := a + b

	// Display result
	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
}
