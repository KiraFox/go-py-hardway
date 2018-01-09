package main

import (
	"fmt"
)

func main() {
	// The ":=" is short hand variable declaration, meaning you can declare a varaible inside a function without using var
	// The Sprintf part of the fmt package allows you to create formatted strings with initalized values and set it to a variable
	x := fmt.Sprintf("There are %d types of people.", 10)
	binary := "binary"
	do_not := "don't"
	y := fmt.Sprintf("Those who know %s and those who %s.", binary, do_not)

	// Print out the variables
	fmt.Println(x)
	fmt.Println(y)

	// Print formatted strings with variable.  Don't forget to use \n to create new line at end of the printed line
	fmt.Printf("I said: %v \n", x)
	fmt.Printf("I also said: '%s'.\n", y)

	// If you set a variable = to true or false, it is a boolean. A boolen can only be true or false. This is explained more later.
	// The formatter for booleans is %t, like %s is for strings.
	hilarious := false
	joke_evaluation := "Isn't that joke so funny?! %t\n"

	// Have to use printf since the var joke_evaluation has a formatter inside it
	// Make sure the var type and the formatter type match, otherwise it won't print correctly.
	fmt.Printf(joke_evaluation, hilarious)

	// Declaring var w & e using shorthand :=
	w := "This is the left side of..."
	e := "a string with a right side."

	// Using the + in the Println 'adds' the 2 variables together in the order they are given, so it creates 1 printed string line
	fmt.Println(w + e)

}
