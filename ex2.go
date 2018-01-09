package main

// A comment, this is so you can read your program later
// Anything after the // is ignored by go if on the same line

import (
	"fmt"
)

func main() {
	fmt.Println("I could have code like this") //and the comment after is ignored
	// You can also use a comment to "disable" or comment out a piece of code
	// fmt.Println("This won't run")
	fmt.Println("This will run")
}

/* This is how
to comment out multiple lines */
