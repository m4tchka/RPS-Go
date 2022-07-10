package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startPrompt() {
	instructions := "Please enter your name"
	fmt.Println(instructions)
}

func main() {
	startPrompt()
	var playerName string
	fmt.Scan(&playerName)
	fmt.Printf("Greetings, %v.", playerName)

}

func rngRoll() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(3)

	switch number {
	case 0:
		fmt.Sprint("Rock")
	case 1:
		fmt.Sprint("Paper")
	default:
		fmt.Sprint("Scissors")
	}

}
