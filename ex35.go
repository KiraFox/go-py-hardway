package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	start()
}

func gold_room() {
	fmt.Println("This room is full of gold. How much do you take?")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	checkError(err)
	choice = strings.TrimSpace(choice)

	how_much, err := strconv.Atoi(choice)
	if err != nil {
		dead("Man, learn to type a number!")
	}

	// This is long way to do it.
	/*var how_much int

	if strings.ContainsAny(choice, "0123456789") {
		how_much, err = strconv.Atoi(choice)
		if err != nil {
			dead("Man, learn to type a number.")
		}
	} else {
		dead("Man, learn to type a number.")
	}*/

	if how_much < 50 {
		fmt.Println("Nice, you're not greedy, you win!")
		os.Exit(0)
	} else {
		dead("You gready bastard!")
	}
}

func bear_room() {
	fmt.Println("There is a bear here.\nThe bear has a bunch of honey.")
	fmt.Println("The fat bear is in front of another door.")
	fmt.Println("How are you going to move the bear?")
	fmt.Print("> ")

	bear_moved := false

	for true {
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		checkError(err)
		choice = strings.TrimSpace(choice)

		if choice == "take honey" {
			dead("The bear looks at you then slaps your face off.")
		} else if choice == "taunt bear" && !bear_moved {
			fmt.Println("The bear has moved from the door. You can go through it now.")
			fmt.Print("> ")
			bear_moved = true
		} else if choice == "taunt bear" && bear_moved {
			dead("The bear gets pissed off and chews your leg off.")
		} else if choice == "open door" && bear_moved {
			gold_room()
		} else {
			fmt.Println("I have no idea what that means.")
		}
	}

}

func cthulhu_room() {
	fmt.Println("Here you see the great evil Cthulhu.")
	fmt.Println("He, it, whatever stares at you and you go insane.")
	fmt.Println("Do you flee for your life or eat your head?")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	checkError(err)
	choice = strings.TrimSpace(choice)

	if strings.Contains(choice, "flee") {
		start()
	} else if strings.Contains(choice, "head") {
		dead("Well taht was tasty!")
	} else {
		cthulhu_room()
	}
}

func dead(why string) {
	fmt.Println(why, "Good job!")
	os.Exit(0)
}

func start() {
	fmt.Println("You are in a dark room.")
	fmt.Println("There is a door to your right and left.")
	fmt.Println("Which one do you take?")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	checkError(err)
	choice = strings.TrimSpace(choice)

	if choice == "left" {
		bear_room()
	} else if choice == "right" {
		cthulhu_room()
	} else {
		dead("You stumble around the room until you starve.")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
