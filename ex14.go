package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	script, user_name, secret := os.Args[0], os.Args[1], os.Args[2]
	promt := "Answer: "
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Hi %s, I'm the %v script.\n", user_name, script)
	fmt.Println("I'd like to ask you a few questions.")

	fmt.Printf("Do you like me %s?\n"+promt, user_name)
	likes, _ := reader.ReadString('\n')
	likes = strings.TrimSpace(likes)

	fmt.Printf("Where do you live %s?\n"+promt, user_name)
	lives, _ := reader.ReadString('\n')
	lives = strings.TrimSpace(lives)

	fmt.Print("What kind of computer do you have?\n" + promt)
	computer, _ := reader.ReadString('\n')
	computer = strings.TrimSpace(computer)

	fmt.Printf(`
Alright, so you said %v about liking me.
You live in %v. Not sure where that is
And you have %v computer. Nice.
And you told me a secret: %v. Oh my.
`, likes, lives, computer, secret)

}
