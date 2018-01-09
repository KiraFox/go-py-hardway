package main

import (
	"fmt"
)

func main() {
	// Variables created inside a function are only seen by that function, no matter if they are capitalized or not
	var (
		my_name   = "Zed A. Shaw"
		my_age    = 35  // not a lie
		my_height = 74  // inches
		my_weight = 180 // lbs
		my_eyes   = "Blue"
		my_teeth  = "White"
		my_hair   = "Brown"
		//Next is a type conversion, where you turn my_height from what it was intialized as (int) into a float64
		//so it can be used mathematically as needed for the var my_height_cm
		my_height_cm = float64(my_height) * 2.54
	)

	// Use Printf instead of Println when making formatted strings.
	// Must use \n before last double quotes to create new line at the end of the string like Println auto does
	// If you don't use \n it prints everything together.

	fmt.Printf("Let's talk about %s.\n", my_name)

	// Use formatter %d for base 10 integers

	fmt.Printf("He's %d inches tall.\n", my_height)
	fmt.Printf("He's %d pounds heavy.\n", my_weight)
	fmt.Println("Actually that's not too heavy.")
	fmt.Printf("He's got %s eyes and %s hair.\n", my_eyes, my_hair)
	fmt.Printf("His teeth are usually %s depending on coffee.\n", my_teeth)

	// This line is tricky, try to get it exactly right
	fmt.Printf("If I add %d, %d, and %d I get %d.\n", my_age, my_height, my_weight, my_age+my_height+my_weight)

	fmt.Println("My height in centimeters is", my_height_cm)
}
