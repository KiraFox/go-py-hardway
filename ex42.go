package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	rover := dog{"Rover"}
	makeItSpeak(rover)
	kira := cat{"Kira"}
	makeItSpeak(kira)

	fmt.Println("Dog is currently: ", rover)

	fileDog, err := os.OpenFile("dog.txt", os.O_RDWR|os.O_CREATE, 0755)
	checkError(err)
	defer fileDog.Close()

	d, err := json.Marshal(rover)
	checkError(err)

	fileDog.Write(d)

	rover = dog{" "}

	fmt.Println("Dog is currently: ", rover)

	text, err := ioutil.ReadFile("dog.txt")
	checkError(err)

	err = json.Unmarshal(text, &rover)
	checkError(err)

	fmt.Println("Dog is currently: ", rover)

}

func makeItSpeak(animal animal) {
	fmt.Printf("%s says: %s\n", animal.name(), animal.speak())
}

type animal interface {
	speak() string
	name() string
}

type dog struct {
	Id string `json:"Id"`
}

func (dog) speak() string {
	return "Bark"
}

func (d dog) name() string {
	return d.Id
}

type cat struct {
	Id string
}

func (cat) speak() string {
	return "Meow"
}

func (c cat) name() string {
	return c.Id
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
