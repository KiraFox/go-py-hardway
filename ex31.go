package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	prompt := "> "
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("You enter a dark room with two doors. Do you go through door 1",
		" or door 2?\n", prompt)

	door, err := reader.ReadString('\n')
	checkError(err)
	door = strings.TrimSpace(door)

	if door == "1" {
		fmt.Println("There's a giant bear eating a cheese cake. What do you do?")
		fmt.Println("1. Take the cake.")
		fmt.Print("2. Scream at the bear.\n", prompt)

		bear, err := reader.ReadString('\n')
		checkError(err)
		bear = strings.TrimSpace(bear)

		if bear == "1" {
			fmt.Println("The bear eats your face off. Good job!")
		} else if bear == "2" {
			fmt.Println("The bear eats your legs off. Good job!")
		} else {
			fmt.Printf("Well, doing %s is probably better. Bear runs away.", bear)
		}
	} else if door == "2" {
		fmt.Println("You stare into the endless abyss at Cthulu's retina.")
		fmt.Println("1. Blueberries.")
		fmt.Println("2. Yellow jacket clothespins.")
		fmt.Print("3. Understanding revolvers yelling melodies.\n", prompt)

		insanity, err := reader.ReadString('\n')
		checkError(err)
		insanity = strings.TrimSpace(insanity)

		// statement OR statement = statement || statement
		// statement AND statement = statement && statement
		// NOT statement = !statement
		if insanity == "1" || insanity == "2" {
			fmt.Println("Your body survives powered by a mind of jello. Good job!")
		} else {
			fmt.Println("The insanity rots your eyes into a pool of muck. Good job!")
		}
	} else {
		fmt.Println("You stumble around and fall on a knife and die. Good job!")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
