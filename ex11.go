package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// set a variable to read the command line iput
	// use the NewReader function from the module bufio
	// tell the reader to look for Standard Input from the OS using os module
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How old are you? ")

	// When you want a response saved from the command line:
	// Create a new variable, _ empty space (this is ignoring possible errors) :=
	// reader varaible using the Read String function that reads until newline(return)
	// Then to get rid of any excess white space before or after the response
	// Use the TrimSpace function from the Strings module and set it = variable
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)

	fmt.Println("How tall are you? ")
	height, _ := reader.ReadString('\n')
	height = strings.TrimSpace(height)

	fmt.Println("How much do you weigh in pounds? ")
	weight, _ := reader.ReadString('\n')
	weight = strings.TrimSpace(weight)

	fmt.Printf("So, you are %s years old, %v tall, and %v lbs heavy.",
		age, height, weight)

}
