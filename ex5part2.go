package main

import (
	"fmt"
)

func main() {
	// Variables created inside a function are only seen by that function, no matter if they are capitalized or not
	var (
		//Making all the #s the same type, float64 so we don't have to use conversion
		my_name      = "Zed A. Shaw"
		my_age       = 35.0  // not a lie
		my_height    = 74.0  // inches
		my_weight    = 180.0 // lbs
		my_eyes      = "Blue"
		my_teeth     = "White"
		my_hair      = "Brown"
		my_height_cm = my_height * 2.54
	)

	// Use Printf instead of Println when making formatted strings.
	// Must use \n before last double quotes to create new line at the end of the string like Println auto does
	// If you don't use \n it prints everything together.

	fmt.Printf("Let's talk about %s.\n", my_name)

	// Use formatter %f for floating point #s.
	// If you want them to only go to a certain amount of decimal places, for example 4 places, use %.4f

	fmt.Printf("He's %.0f inches tall.\n", my_height)
	fmt.Printf("He's %.0f pounds heavy.\n", my_weight)
	fmt.Println("Actually that's not too heavy.")
	fmt.Printf("He's got %s eyes and %s hair.\n", my_eyes, my_hair)
	fmt.Printf("His teeth are usually %s depending on coffee.\n", my_teeth)

	// This line is tricky, try to get it exactly right
	fmt.Printf("If I add %.0f, %.0f, and %.0f I get %.0f.\n", my_age, my_height, my_weight, my_age+my_height+my_weight)

	fmt.Println("My height in centimeters is", my_height_cm)
}
