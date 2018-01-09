package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	fmt.Printf(`
We're going to erase %v.
If you don't want that, hit CTRL-C (^C).
If you do want that, hit RETURN.
`, filename)

	// Now we are going to start checking for errors using a function we create
	// See below for the check error function.
	// We are starting this to get used to Go Code Best Practices.
	fmt.Print("?")
	reader := bufio.NewReader(os.Stdin)
	_, err = reader.ReadString('\n')
	checkError(err)

	// OpenFile is needed for granting Write permission for the file you open
	// You need a string with the file name "test.txt" or a variable set to that
	// Then you need to set the Flag (Read, Read/Write, Create, etc):
	// For Flag information, look at the Constants section in OS package documentation
	// Here we set it to Read/Write|Create, which means Open the File and give
	// Read/Write permission OR Create the file if none exists with Read/Write permission
	// Next you need to set Linux File Permission code# so the program can dictate
	// Who has what kind of access to that file later.
	fmt.Println("Opening the file...")
	target, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	checkError(err)
	defer target.Close()

	// Here we use os.Truncate to delete all of the file contents.
	// When using os.Truncate you need a string with the file name "test.txt",
	// or a variable set to string of the file name.
	// Then you set how many bytes you want the file to end up having.
	// If we wanted the file to be 100 bytes, and it was originally more than 100
	// this will delete anything past the 100th byte.  If it originally had less
	// than 100 bytes, it would keep the original contents and fill in difference
	// with null bytes until it reached 100 bytes total.
	// In this case, we want to delete all the contents so we set parameter to 0
	fmt.Println("Truncating the file. Goodbye!")
	err = os.Truncate(filename, 0)
	checkError(err)

	fmt.Print("Now I'm going to ask you for three lines.\nLine 1: ")
	line1, err := reader.ReadString('\n')
	checkError(err)

	fmt.Print("Line 2: ")
	line2, err := reader.ReadString('\n')
	checkError(err)

	fmt.Print("Line 3: ")
	line3, err := reader.ReadString('\n')
	checkError(err)

	_, err = target.WriteString(line1 + line2 + line3)
	checkError(err)

}

// If you want to check for errors and only write out once what you want to
// happen if an error does occur, then you can create another function to call
// when you check for errors.  This will save time and space when coding.
// First we have func then the name of the function (whatever we want to call it)
// Then we have (parameter type) which is like creating a variable and giving it
// a type. Then you can tell it to return something or leave it blank if you
// do not want something at the end. Then {what you want to happen in this function}.
// Here we are saying we want a function named checkError(with parameter called
// err with a type of error) to return nothing {but tell us if there is an error
// to log it and output the error as a string in the terminal and exit}
// Anytime you want to check for errors now, you can call the function as shown above.
// Eventually you will need to do different things for different errors because
// you don't always want the the program to stop running when it encounters an
// error during production.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
