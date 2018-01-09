package main

import (
	"fmt"
)

func main() {
	// This is a slice
	// Can append only to slices
	the_count := []int{1, 2, 3, 4, 5}

	// This is an array.
	// Arrays can only be the size that is specified in the []
	// THis array can only ever have 5 pieces of data in it.
	// the_count := [5]int{1, 2, 3, 4, 5}

	fruits := []string{"apples", "oranges", "pears", "apricots"}

	// "interface{}" means any type
	change := []interface{}{1, "pennies", 2, "dimes", 3, "quarters"}

	// for = begins for-loop and is treated like a function. runs until condition is completed or False
	// _(index variable name, use _ to ignore this returned value),
	// number(value variable name to hold the returned info) :=
	// range(returns index# to index variable if there is one, and returns info
	// at that index# to the value variable, and makes for-loop run through whole list)
	// the_count(name of array or slice you want range to return values of) {
	// (what you want done each time this for-loop runs)
	// Print out the value assigned to the varaible "number"}
	// **The index variable and the value variable are local only to this for-loop
	for _, number := range the_count {
		fmt.Printf("This is count %d\n", number)
	}

	// for = begins for-loop
	// index (index variable name) := 0 (start the for-loop at this index location);
	// index < len(fruits) *run for-loop until the index# is more than the length
	// of the array/slice "fruits" making the conidition False* ;
	// index++ (tells for loop to increase index by 1 each iteration, same as index+=1) {
	// (what you want done each time the for-loop runs)
	// Print out info in the array/slice "fruits" at [index#]}
	for index := 0; index < len(fruits); index++ {
		fmt.Printf("This is count %s\n", fruits[index])
	}

	// Remember to use the formatter %v to get the raw data if type is unknown or changes
	// _ (ignored index variable) ; x (value variable)
	for _, x := range change {
		fmt.Printf("I have %v\n", x)
	}

	// This creates a variable "elements" set to an empty slice of integers with no
	// predetermined length
	var elements []int

	for i := 0; i < 6; i++ {
		fmt.Printf("Adding %d to the list.\n", i)
		elements = append(elements, i)
	}

	fmt.Println(elements)

	// You can also "make" an empty slice with an intiliazed length and capacity
	// You can always add on to the length and capacity as needed

	colors := make([]string, 1, 1)

	fmt.Println("This is currently what is inside colors")
	fmt.Println(colors)
	fmt.Println("This is the current length of colors: ", len(colors))

	fmt.Println("Let's add a color so the slice is no longer empty.")
	colors[0] = "brown"

	fmt.Println("Color now: ", colors)

	fmt.Println("Let's add some new colors to our current colors.")

	new_colors := []string{"blue", "black", "purple", "green"}

	colors = append(colors, new_colors...)

	fmt.Println("Now the new length of colors is:", len(colors))
	fmt.Println("And the contents of colors: ", colors)

	fmt.Println("Let's change that first color.")
	colors[0] = "fuschia"

	fmt.Println(colors)
}

/*
	var totalTime time.Duration
	var prevTime time.Time
	var prevIsIn bool

	for scanner.Scan() {
		text := scanner.Text()
		currIsIn, currTime, err := parseEntry(text)


		if prevIsIn && !currIsIn {

		}

		prevTime = currTime
		prevIsIn = currIsIn
	}*/
