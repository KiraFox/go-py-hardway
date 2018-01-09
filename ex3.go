// Package main signifies this is an executable file
package main

// Imports packages that give us usable commands
import (
	"fmt"
)

// Main funcation that the package executes when run
func main() {
	// Outputs the line
	fmt.Println("I will now count my chickens:")

	// Outputs the world and final math result
	fmt.Println("Hens", 25+30/6)
	fmt.Println("Roosters", 100-25*3%4)

	fmt.Println("Now I will count the eggs:")

	fmt.Println(3 + 2 + 1 - 5 + 4%2 - 1/4 + 6)

	fmt.Println("Is it true that 3 + 2 < 5 - 7")

	fmt.Println(3+2 < 5-7)

	fmt.Println("What is 3 + 2?", 3+2)
	fmt.Println("What is 5 - 7?", 5-7)

	fmt.Println("Oh, that's why it's False.")

	fmt.Println("How about some more.")

	fmt.Println("Is it greater?", 5 > -2)
	fmt.Println("Is it greater or equal?", 5 >= -2)
	fmt.Println("Is it less or equal?", 5 <= -2)

	// Floating point numbers means numbers with decimal points.
	fmt.Println("Trying floating point numbers", 25.0/3.0)
}
