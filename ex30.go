package main

import (
	"fmt"
)

func main() {
	people := 30
	cars := 40
	trucks := 15

	// if condition { } else if condition { } else { }
	// Whenever using else if and else statements, make sure they begin on the
	// same line as the closing curly bracket } of the preceding statment.
	// This tells Go that these statements go together and to run that block of
	// code until the one of the statements is true, otherwise it will run the
	// else statement
	if cars > people {
		fmt.Println("We should take the cars.")
	} else if cars < people {
		fmt.Println("We should not take the cars.")
	} else {
		fmt.Println("We can't decide.")
	}

	if trucks > cars {
		fmt.Println("That's too many trucks.")
	} else if trucks < cars {
		fmt.Println("Maybe we should take the trucks.")
	} else {
		fmt.Println("We still can't decide.")
	}

	if people > trucks {
		fmt.Println("Alright, let's just take the trucks.")
	} else {
		fmt.Println("Fine, let's stay home then.")
	}

}
