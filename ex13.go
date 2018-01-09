package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args is what is called an array
	// You must dictate what position in the array you want for each var
	// When running the script in command line:
	// you must put in the same # of arguments as vars you declared after
	// the script name
	// The name of the script will always be position 0
	script, first, second, third := os.Args[0], os.Args[1],
		os.Args[2], os.Args[3]

	fmt.Println("This is the script:", script)
	fmt.Println("This is the first variable:", first)
	fmt.Println("This is the second variable:", second)
	fmt.Println("This is the third variable:", third)

	// Now to add in asking for further input during the script run

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input another variable: ")
	fourth, _ := reader.ReadString('\n')
	fourth = strings.TrimSpace(fourth)

	fmt.Println("You input: ", fourth)
}
