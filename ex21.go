package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's do some math with just functions!")

	age := add(30, 5)
	height := subtract(78, 4)
	weight := multiply(90, 2)
	iq := divide(100, 2)

	fmt.Printf("Age: %d, Height: %d, Weight: %d, IQ: %d\n", age, height, weight, iq)

	// A puzzle for extra credit, type it anyway.
	fmt.Println("Here is a puzzle.")

	what := add(age, subtract(height, multiply(weight, divide(iq, 2))))

	fmt.Println("That becomes: ", what, "Can you do it by hand?")
}

func add(a, b int) int {
	fmt.Printf("ADDING %d + %d\n", a, b)
	return a + b
}

func subtract(a, b int) int {
	fmt.Printf("SUBTRACTING %d - %d\n", a, b)
	return a - b
}

func multiply(a, b int) int {
	fmt.Printf("MULTIPLYING %d * %d\n", a, b)
	return a * b
}

func divide(a, b int) int {
	fmt.Printf("DIVIDING %d / %d\n", a, b)
	return a / b
}
