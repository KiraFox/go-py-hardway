package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("spike.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	f.WriteString("test test")

}
