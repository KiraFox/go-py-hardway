package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Order is important when you are running scripts. Go runs the first line
	// first, then the second, and so on whenever you run your main function.
	// It can pull other functions listed, but it does it in the order the main
	// one tells them to pull.

	input_file := os.Args[1]

	// We have to open the file first for when we tell a function to run on it
	current_file, err := os.Open(input_file)
	checkError(err)
	// Defer the closing until the main function has finished running
	defer current_file.Close()

	fmt.Println("First let's print the whole file:\n")

	print_all(current_file)

	fmt.Println("\nNow let's rewind, kind of like a tape.\n")

	rewind(current_file)

	// We call for a NewScanner after we have done a Seek (rewind function)
	// because Scanners have their own internal track of where they last scanned
	// and Seek does not reset a Scanner.  If we had to Seek again after these
	// Scans then we need call another NewScanner.
	scanner := bufio.NewScanner(current_file)

	fmt.Println("Let's print 3 lines:")

	current_line := 1

	print_a_line(current_line, scanner)

	// This is the same as writing current_line = current_line + 1
	current_line += 1
	print_a_line(current_line, scanner)

	// This is the same as writing current_line = current_line + 1
	current_line++
	print_a_line(current_line, scanner)
}

func print_all(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func rewind(f *os.File) {
	// The first number after Seek is the offset (how many bytes to move) and it
	// can be positive or negative.
	// The second number is whence (from where) you want the offset to go.
	// 0 = Beginning of file
	// 1 = Current position
	// 2 = End of file
	// So here (0,0) means move 0 bytes from the begining of the file.
	_, err := f.Seek(0, 0)
	checkError(err)
}

// We can't have a NewScanner inside this function because each time the function
// would be called, the Scanner would lose its place in the file and be reset.
func print_a_line(line_count int, scanner *bufio.Scanner) {
	scanner.Scan()
	fmt.Println(line_count, scanner.Text())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
