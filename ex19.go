package main

import (
	"fmt"
)

func main() {
	fmt.Println("We can just give the function numbers directly:")
	cheese_and_crackers(20, 30)

	fmt.Println("OR, we can use variables from our script:")
	amount_of_cheese := 10
	amount_of_crackers := 50
	cheese_and_crackers(amount_of_cheese, amount_of_crackers)

	fmt.Println("We can even do math inside too:")
	cheese_and_crackers(10+20, 5+6)

	fmt.Println("And we can combine the two, variables and math:")
	cheese_and_crackers(amount_of_cheese+100, amount_of_crackers+1000)
}

func cheese_and_crackers(cheese_count, boxes_of_crackers int) {
	fmt.Printf("You have %d cheeses!\n", cheese_count)
	fmt.Printf("You have %d boxes of crackers!\n", boxes_of_crackers)
	fmt.Println("Man that's enough for a pary!")
	fmt.Println("Get a blanket.\n")
}
