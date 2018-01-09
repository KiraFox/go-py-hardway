package main

import (
	"fmt"
)

func main() {
	var happyBDay Song

	happyBDay.sing("Happy birthday to you.",
		"I don't want to get sued",
		"So I'll stop there")

	var bullsParade Song

	bullsParade.sing("They rally around tha family",
		"With pockets full of shells.")
}

type Song struct {
	lyrics string
}

func (self *Song) sing(line ...string) {
	for i, _ := range line {
		self.lyrics = line[i]
		fmt.Println(self.lyrics)
	}
}
