package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	// Use ReadFile only if you know file is small
	// Use the _ instead of err if you want to ignore any errors
	txt, _ := ioutil.ReadFile(filename)

	fmt.Printf("Here's your file %v:\n", filename)
	fmt.Println(string(txt))

	fmt.Print("Type the filename again:\n> ")

	reader := bufio.NewReader(os.Stdin)
	file_again, _ := reader.ReadString('\n')
	file_again = strings.TrimSpace(file_again)

	// Here we are going to start checking for Errors after something happens
	// Use the OS package for opening/using files of unknown size
	// os.Open only opens in READONLY
	// If an error happens while trying to Open the file,
	// the error will be shown in the terminal
	// defer means X will happen when the function it's inside of is done running
	// In this case it means it will Close the file once the function ends
	txt_again, err := os.Open(file_again)
	if err != nil {
		log.Fatal(err)
	}
	defer txt_again.Close()

	// Here we are going Scan (read) the file we opened.
	// Just like reader, we are creating the variable scanner and
	// setting it to a new buffered scan with the parameter we want scanned.
	// In this case, the parameter is the file we opened.
	// The buffer scan will only scan 1 line at a time whenever you call scanner.scan
	// This means you only use memory the size of the line you scanned and
	// it reuses the memory each time it is called.
	// In order to Scan the whole file and print out the contents in strings,
	// we have to use a 'for loop' to constantly call the scanner on the file.
	// We tell it for every time scanner.Scan is called (which it does in a loop),
	// print out the scanned lines in Text(strings).
	// The for loop will keep running until scanner.Scan has nothing to scan,
	// which means there are no new lines to scan in the file, so therefor
	// you have reached the end of the file.
	scanner := bufio.NewScanner(txt_again)
	fmt.Println("Here is the file again:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
