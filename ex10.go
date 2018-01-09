package main

import (
	"fmt"
)

func main() {
	tabby_cat := "\tI'm tabbed in."
	persian_cat := "I'm split\non a line."
	backslash_cat := "I'm \\ a \\ cat."

	// Cannot use escape sequences in multiline that use ` (back tic)
	// Use the + sign to combine multiple strings in a single var
	// This will allow you to keep your code looking clean still
	fat_cat := "I'll do a list: \n\t* Cat Food\n\t* Fishies\n\t" +
		"* Catnip\n\t* Grass"

	// If using ` (back tic) for multie line strings
	// Type exactly what you want on the lines between tics
	big_cat := `
This is left aligned.
	This is tabbed in.
	* Tabbed list first line
	* Second line
And back again.
`
	fmt.Println(tabby_cat)
	fmt.Println(persian_cat)
	fmt.Println(backslash_cat)
	fmt.Println(fat_cat)
	fmt.Println(big_cat)
}
