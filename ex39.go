package main

import (
	"fmt"
)

func main() {
	// create a mapping of state to abbreviation
	states := map[string]string{
		"Oregon":     "OR",
		"Florida":    "FL",
		"California": "CA",
		"New York":   "NY",
		"Michigan":   "MI"}

	// create a mapping of states and some cities in the,
	cities := map[string]string{
		"CA": "San Francisco",
		"MI": "Detroit",
		"FL": "Jacksonville"}

	// add somce more cities
	cities["NY"] = "New York"
	cities["OR"] = "Portland"

	dash := "----------"

	// print out some cities
	fmt.Println(dash)
	fmt.Println("NY State has: ", cities["NY"])
	fmt.Println("OR State has: ", cities["OR"])

	// print out some states
	fmt.Println(dash)
	fmt.Println("Michigan's abbreviation is: ", states["Michigan"])
	fmt.Println("Florida's abbreviation is: ", states["Florida"])

	// print cities using states
	fmt.Println(dash)
	fmt.Println("Michigan has: ", cities[states["Michigan"]])
	fmt.Println("Florida has: ", cities[states["Florida"]])

	// print every state abbreviation
	fmt.Println(dash)
	for state, abbrev := range states {
		fmt.Printf("%s is abbreviated %s.\n", state, abbrev)
	}

	// print every city in state
	fmt.Println(dash)
	for abbrev, city := range cities {
		fmt.Printf("%s has the city %s.\n", abbrev, city)
	}

	// now do both at the same time
	fmt.Println(dash)
	for state, abbrev := range states {
		fmt.Printf("%s state is abbreviated %s and has city %s.\n", state, abbrev,
			cities[abbrev])
	}
	fmt.Println(dash)

	// A two-value assignment tests for the existence of a key.
	// state, ok := states["Texas"]
	// In this statement, the first value (state) is assigned the value stored under
	// the key "Texas". If that key doesn't exist, state is the value type's zero
	// value (" "). The second value (ok) is a bool that is true if the key exists
	// in the map, and false if not.
	// To test for a key without retrieving the value,
	// use an underscore in place of the first value:

	// safely get an abbreviation by state that might not be there
	_, ok := states["Texas"]
	if !ok {
		fmt.Println("Sorry, no Texas.")
	} else {
		state := states["Texas"]
		fmt.Println(state)
	}

	// get a city with a default value (in this case, empty string)
	city, ok := cities["TX"]
	fmt.Printf("The city for the state 'TX' is: %s.\n", city)
}
