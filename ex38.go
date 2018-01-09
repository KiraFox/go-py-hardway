package main

import (
	"fmt"
	"strings"
)

func main() {
	ten_things := "Apples Oranges Crows Telephone Light Sugar"

	fmt.Println("Wait there are not 10 things in that list. Let's fix that.")

	stuff := strings.Split(ten_things, " ")

	more_stuff := []string{"Day", "Night", "Song", "Frisbee", "Corn", "Banana"}

	for len(stuff) != 10 {
		next_one := more_stuff[0]
		more_stuff = more_stuff[1:]
		fmt.Println("Adding: ", next_one)
		stuff = append(stuff, next_one)
		fmt.Printf("There are %d items now.\n", len(stuff))
	}

	fmt.Println("There we go: ", stuff)

	fmt.Println(stuff[1])
	fmt.Println(stuff[len(stuff)-1])
	fmt.Println(strings.Join(stuff, " "))
	fmt.Println(strings.Join(stuff[3:5], "#"))
}
