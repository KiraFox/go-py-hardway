package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	from_file, to_file := os.Args[1], os.Args[2]

	fmt.Printf("Copying from %s to %s.\n", from_file, to_file)

	info, err := os.Stat(from_file)
	checkError(err)

	fmt.Printf("The input file is %d bytes long.\n", info.Size())

	fmt.Println("File has been copied.")

	indata, err := ioutil.ReadFile(from_file)
	checkError(err)

	out_file, err := os.OpenFile(to_file, os.O_RDWR|os.O_CREATE, 0755)
	checkError(err)
	defer out_file.Close()

	out_file.Write(indata)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
