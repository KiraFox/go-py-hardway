package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's practice everything.")
	fmt.Println("You'd need to know 'bout escapes with \\ that do \n newlines and\t tabs.")

	poem := `
	The lovely world
with logic so firmly planted
cannot discern
 the needs of love
nor comprehend passion from intuition
and requires an explanation
	where there is none.
`

	fmt.Println("----------")
	fmt.Println(poem)
	fmt.Println("----------")

	five := 10 - 2 + 3 - 6
	fmt.Printf("This should be five: %d\n", five)

	start_point := 10000
	beans, jars, crates := secret_formula(start_point)

	fmt.Printf("With a starting point of: %d\n", start_point)
	fmt.Printf("We'd have %d beans, %d jars, and %d crates.", beans, jars, crates)

	start_point = start_point / 10

	// YOu cannot do it the following way unless you are return multiple slices.
	//fmt.Println("We can also do that this way:")
	//fmt.Printf("We'd have %d beans, %d jars, and %d crates", secret_formula(start_point))
}

func secret_formula(started int) (int, int, int) {
	jelly_beans := started * 500
	jars := jelly_beans / 10000
	crates := jars / 100
	return jelly_beans, jars, crates
}
