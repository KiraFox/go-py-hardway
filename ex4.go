package main

import (
	"fmt"
)

// If variables are going to be interacting with each other they must be the same type.
// You cannot have one var be type int (whole #s) and another type float64 (floating point #s) if you want them to interact
// Variables listed outside funcs are global/viewable to other files starting with the same package name
// Variable name starting with capital letters are 'exported' (other package name files can see/use them)
var (
	cars                       = 100.0
	space_in_a_car             = 4.0
	drivers                    = 30.0
	passengers                 = 90.0
	cars_not_driven            = cars - drivers
	cars_driven                = drivers
	carpool_capacity           = cars_driven * space_in_a_car
	average_passengers_per_car = passengers / cars_driven
)

func main() {
	fmt.Println("There are", cars, "cars available.")
	fmt.Println("There are only", drivers, "drivers available.")
	fmt.Println("There will be", cars_not_driven, "empty cars today.")
	fmt.Println("We can transport", carpool_capacity, "people today.")
	fmt.Println("We have", passengers, "to carpool today.")
	fmt.Println("We need to put about", average_passengers_per_car, "in each car.")
}
