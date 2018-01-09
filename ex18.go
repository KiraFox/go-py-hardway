package main

import (
	"fmt"
)

func main() {

	print_two("Sekuria", "Fox")
	print_two_again("Hello", "Goodbye")
	print_one("First")
	print_none()
}

func print_two(args ...string) {
	arg1, arg2 := args[0], args[1]
	fmt.Printf("arg1: %v, arg2: %v\n", arg1, arg2)
}

func print_two_again(arg1, arg2 string) {
	fmt.Printf("arg1: %v, arg2: %v\n", arg1, arg2)
}

func print_one(arg1 string) {
	fmt.Printf("arg1: %v\n", arg1)
}

func print_none() {
	fmt.Println("I got nothin'.")
}
