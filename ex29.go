package main

import (
	"fmt"
)

func main() {
	people := 20
	cats := 30
	dogs := 15

	if people < cats {
		fmt.Println("Too many cats! The world is doomed!")
	}

	if people > cats {
		fmt.Println("Not too many cats! The world is saved!")
	}

	if people < dogs {
		fmt.Println("The world is drooled on!")
	}

	if people > dogs {
		fmt.Println("The world is dry!")
	}

	dogs += 5

	if people >= dogs {
		fmt.Println("People are greater than or equal to dogs.")
	}

	if people <= dogs {
		fmt.Println("People are less than or equal to dogs.")
	}

	if people == dogs {
		fmt.Println("People are dogs.")
	}
}
