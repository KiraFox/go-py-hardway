package main

import (
	"fmt"
)

func main() {
	// Print single line of the string
	fmt.Println("Mary had a little lamb.")
	// Print formatted string with new line "\n" at end
	fmt.Printf("Its fleece was white as %s\n.", "snow")
	fmt.Println("And everywhere that Mary went.")

	// Create variable s as an empty string
	s := ""
	// This is a 'for' loop that will execute 10 times
	for i := 0; i < 10; i++ {
		// We are making variable s = . then adding it to itself and making it equal to the new value,
		// for as many times as the loop runs
		s += "."
	}

	// Print out what the new value of s is after being put through the for loop
	fmt.Println(s)

	// Intializing variables
	var (
		end1  = "C"
		end2  = "h"
		end3  = "e"
		end4  = "e"
		end5  = "s"
		end6  = "e"
		end7  = "B"
		end8  = "u"
		end9  = "r"
		end10 = "g"
		end11 = "e"
		end12 = "r"
	)

	// Print out the variables values without creating new line at end
	fmt.Print(end1 + end2 + end3 + end4 + end5 + end6)
	// Continues on same line as previous print and adds a space at the end of it
	fmt.Print(" ")
	// Continues on the same line as previous print and adds the values of the varaiables.
	fmt.Println(end7 + end8 + end9 + end10 + end11 + end12)
}
