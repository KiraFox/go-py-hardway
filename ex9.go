// Here's some new strange stuff, remember type it exactly.
package main

import (
	"fmt"
)

func main() {
	days := "Mon Tue Wed Thu Fri Sat Sun"
	months := "Jan\nFeb\nMar\nApr\nMay\nJun\nJul\nAug"

	fmt.Println("Here are the days: ", days)
	fmt.Println("Here are the months: ", months)

	// This format has the lines all left aligned when printed
	fmt.Println(`There is something going on here.
With the single back-tick.
We'll be able to type as much as we like.
Even 4 lines if we want, or 5, or 6.`)

	// Printed lines 2+ will be show indented 2 times
	fmt.Println(`This type of format
		Will indent 2nd and later lines 2 times.`)
}
